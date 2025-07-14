package handles

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Muskan2518/fealtyX/models"
)

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updated models.Student
	json.NewDecoder(r.Body).Decode(&updated)

	mu.Lock()
	defer mu.Unlock()
	s, exists := models.Students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	if updated.Name != "" {
		s.Name = updated.Name
	}
	if updated.Age > 0 {
		s.Age = updated.Age
	}
	if updated.Email != "" {
		s.Email = updated.Email
	}

	models.Students[id] = s
	json.NewEncoder(w).Encode(s)
}
