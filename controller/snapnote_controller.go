package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stenstromen/go-snapnote-backend/model"
	"github.com/stenstromen/go-snapnote-backend/service"
)

func CreateFormData(w http.ResponseWriter, r *http.Request) {
	var formData model.FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var noteID string
	noteID, err = service.InsertFormData(formData)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, noteID)
}

func GetFormData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID := vars["noteid"]

	formData, err := service.GetFormDataByNoteID(noteID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Note not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		}
		return
	}

	jsonData, err := json.Marshal(formData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
