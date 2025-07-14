package handles

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Muskan2518/fealtyX/models"
)

var mu sync.Mutex

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	json.NewDecoder(r.Body).Decode(&s)

	if s.Name == "" || s.Age <= 0 || s.Email == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	s.ID = len(models.Students) + 1
	models.Students[s.ID] = s
	json.NewEncoder(w).Encode(s)
}
