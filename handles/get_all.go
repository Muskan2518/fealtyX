package handles

import (
	"encoding/json"
	"net/http"

"github.com/Muskan2518/fealtyX/models"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var result []models.Student
	for _, s := range models.Students {
		result = append(result, s)
	}
	json.NewEncoder(w).Encode(result)
}
