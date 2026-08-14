package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"strconv"
	"strings"
	"sync"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

func init() {
	teachers[nextID] = models.Teacher{ID: nextID, FirstName: "John", LastName: "Doe", Subject: "Math", Class: "10A"}
	nextID++
	teachers[nextID] = models.Teacher{ID: nextID, FirstName: "Jane", LastName: "Smith", Subject: "Science", Class: "10B"}
	nextID++
	teachers[nextID] = models.Teacher{ID: nextID, FirstName: "Jane", LastName: "Doe", Subject: "Fishing", Class: "10C"}
	nextID++
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println(idStr)

	if idStr == "" {

		firstName := r.URL.Query().Get("firstName")
		lastName := r.URL.Query().Get("lastName")
		teacherList := make([]models.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if firstName != "" && teacher.FirstName != firstName {
				continue
			}
			if lastName != "" && teacher.LastName != lastName {
				continue
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
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}

	// Handle the case where an ID is provided
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

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

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST request to add a new teacher
	mutex.Lock()
	defer mutex.Unlock()
	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]models.Teacher, 0, len(newTeachers))

	for _, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers = append(addedTeachers, newTeacher)
		nextID++
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

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	//teachers/{id}
	fmt.Println(r.Method)
	msg := fmt.Sprintf("Hello %s Teachers Route", r.Method)
	fmt.Println(msg)
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
	case http.MethodPost:
		addTeacherHandler(w, r)
	case http.MethodPut:
		w.Write([]byte(msg))
		fmt.Println(msg)
	case http.MethodPatch:
		w.Write([]byte(msg))
		fmt.Println(msg)
	case http.MethodDelete:
		w.Write([]byte(msg))
		fmt.Println(msg)
	}
}
