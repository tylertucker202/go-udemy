package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
)

func GetTeacherHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	fmt.Println(idStr)

	teacher, err := sqlconnect.GetTeacherByID(idStr)
	if err != nil {
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

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var teachers []models.Teacher
	idStr := r.PathValue("id")

	fmt.Println(idStr)
	teachers, err := sqlconnect.GetTeachersByID(r)
	if err != nil {
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teachers),
		Data:   teachers,
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

	addedTeachers, err := sqlconnect.AddTeacherDBHandler(r)
	if err != nil {
		return
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

func UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {

	updatedTeacher, err := sqlconnect.UpdateTeacherDBHandler(r)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "applcation/json")
	json.NewEncoder(w).Encode(updatedTeacher)

}

func PatchTeachersHandler(w http.ResponseWriter, r *http.Request) {

	err := sqlconnect.PatchTeachersDBHandler(r)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func PatchTeacherHandler(w http.ResponseWriter, r *http.Request) {

	teacher, err := sqlconnect.PatchTeacherDBHandler(r)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "applcation/json")
	json.NewEncoder(w).Encode(teacher)

}

func DeleteTeachersHandler(w http.ResponseWriter, r *http.Request) {

	deletedIds, err := sqlconnect.DeleteTeachersDBHandler(r)
	if err != nil {
		return
	}

	// Response Body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		IDs    []int  `json:"ids"`
	}{
		Status: "Teachers sucessfully Deleted",
		IDs:    deletedIds,
	}

	json.NewEncoder(w).Encode(response)

}

func DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	id, err := sqlconnect.DeleteTeacherDBHandler(r)
	if err != nil {
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
