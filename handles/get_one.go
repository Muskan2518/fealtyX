package handles

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Muskan2518/fealtyX/models"
	
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	defer mu.Unlock()
	s, exists := models.Students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(s)
}
