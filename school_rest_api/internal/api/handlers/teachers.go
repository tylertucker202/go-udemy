package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

func isValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidSortField(field string) bool {
	validFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"class":      true,
		"subject":    true,
	}
	return validFields[field]
}
func GetTeacherHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	fmt.Println(idStr)

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	// Handle the case where an ID is provided
	var teacher models.Teacher

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(
		&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database query err", http.StatusInternalServerError)
	}
	defer db.Close()

	response := struct {
		Status string         `json:"status"`
		Data   models.Teacher `json:"data"`
	}{
		Status: "success",
		Data:   teacher,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	fmt.Println(idStr)

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
	var args []interface{}
	query, args = addFilters(r, query, args)

	sortParams := r.URL.Query()["sortby"]

	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}
			field, order := parts[0], parts[1]
			if !isValidSortOrder(order) || !isValidSortField(field) {
				continue
			}
			if i > 0 {
				query += ","
			}
			query += " " + field + " " + order
		}
	}

	fmt.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database Query Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	teacherList := make([]models.Teacher, 0)
	for rows.Next() {
		var teacher models.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			http.Error(w, "Error scanning database results", http.StatusInternalServerError)
		}
		teacherList = append(teacherList, teacher)

	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teacherList),
		Data:   teacherList,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST request to add a new teacher

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?,?,?,?,?)")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in preparing SQL query: %s", err), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var newTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			http.Error(w, "Error inserting data into the database", http.StatusBadRequest)
			return
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting last insert ID", http.StatusBadRequest)
			return
		}
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func addFilters(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}
	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += fmt.Sprintf(" AND %s = ?", dbField)
			args = append(args, value)
		}
	}
	return query, args
}

func UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTeacher models.Teacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, subject, class from teachers WHERE id = ?", id).Scan(
		&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			http.Error(w, "Teacher not found", http.StatusNotFound)
		}
		log.Println(err)
		http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return
	}

	updatedTeacher.ID = teacher.ID
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? where id = ?",
		updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID,
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error updating teacher", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "applcation/json")
	json.NewEncoder(w).Encode(updatedTeacher)

}

func PatchTeachersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var updates []map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return
	}

	fmt.Println(updates)

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("invalid teacher id %s in update", update["id"]), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("invalid teacher id %s in update", update["id"]), http.StatusBadRequest)
			return
		}

		var teacherFromDb models.Teacher
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ?", id).Scan(
			&teacherFromDb.ID, &teacherFromDb.FirstName, &teacherFromDb.LastName, &teacherFromDb.Email, &teacherFromDb.Class, &teacherFromDb.Subject,
		)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				log.Println(err)
				http.Error(w, "Teacher not found", http.StatusNotFound)
			}
			log.Println(err)
			http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
			return
		}

		teacherVal := reflect.ValueOf(&teacherFromDb).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" {
				continue
			}
			for i := 0; i < teacherVal.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("cannot convert %v to %v", val.Type(), fieldVal.Type())
						}
					}
				}
			}
		}

		_, err = tx.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?",
			teacherFromDb.FirstName, teacherFromDb.LastName, teacherFromDb.Email, teacherFromDb.Class, teacherFromDb.Subject, teacherFromDb.ID)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}

func PatchTeacherHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, subject, class from teachers WHERE id = ?", id).Scan(
		&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			http.Error(w, "Teacher not found", http.StatusNotFound)
		}
		log.Println(err)
		http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return
	}
	// Apply updates using reflect
	teacherVal := reflect.ValueOf(&teacher).Elem()
	teacherType := teacherVal.Type()

	for k, v := range updates {
		for i := 0; i < teacherVal.NumField(); i++ {
			field := teacherType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				if teacherVal.Field(i).CanSet() {
					fieldVal := teacherVal.Field(i)
					valType := teacherVal.Field(i).Type()
					fieldVal.Set(reflect.ValueOf(v).Convert(valType))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? where id = ?",
		teacher.FirstName, teacher.LastName, teacher.Email, teacher.Class, teacher.Subject, teacher.ID,
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error updating teacher", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "applcation/json")
	json.NewEncoder(w).Encode(teacher)

}

func DeleteTeachersHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Database deletion error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error retrieving delete result", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Teacher Not Found", http.StatusNotFound)
		return
	}

	// Response Body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Teacher sucessfully Deleted",
		ID:     id,
	}

	json.NewEncoder(w).Encode(response)

}
