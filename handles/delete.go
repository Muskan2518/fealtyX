package handles

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Muskan2518/fealtyX/models"

)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	defer mu.Unlock()
	if _, exists := models.Students[id]; !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	delete(models.Students, id)
	w.WriteHeader(http.StatusNoContent)
}
