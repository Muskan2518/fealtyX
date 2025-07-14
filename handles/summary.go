package handles

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/Muskan2518/fealtyX/models"
)

func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	// Lock to access the student map
	mu.Lock()
	student, exists := models.Students[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	// Prepare prompt for Ollama
	prompt := fmt.Sprintf("Summarize this student:\nName: %s Age: %d Email: %s", student.Name, student.Age, student.Email)

	// Send request to Ollama
	reqBody, _ := json.Marshal(map[string]string{
		"model":  "gemma:2b",
		"prompt": prompt,
	})

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to contact Ollama", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read streaming response from Ollama
	var summaryBuilder strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		var chunk map[string]interface{}
		if err := json.Unmarshal([]byte(line), &chunk); err == nil {
			if piece, ok := chunk["response"].(string); ok {
				summaryBuilder.WriteString(piece)
			}
		}
	}

	// Clean the raw summary text
	rawSummary := summaryBuilder.String()

	// Remove common intro phrases
	rawSummary = strings.TrimPrefix(rawSummary, "Sure. Here's a summary of the student:")
	rawSummary = strings.TrimPrefix(rawSummary, "Here is a summary of the student:")
	rawSummary = strings.TrimPrefix(rawSummary, "Summary:")

	// Replace newlines, bullets, and multiple spaces
	clean := strings.ReplaceAll(rawSummary, "\n", " ")
	clean = strings.ReplaceAll(clean, "*", "")
	clean = strings.ReplaceAll(clean, "•", "")
	clean = strings.TrimSpace(clean)
	clean = strings.Join(strings.Fields(clean), " ") // remove extra spaces

	// Format: "Name: John, Age: 21, Email: x"
	clean = strings.ReplaceAll(clean, " Name:", "Name:")
	clean = strings.ReplaceAll(clean, " Age:", ", Age:")
	clean = strings.ReplaceAll(clean, " Email:", ", Email:")

	// Send final cleaned response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"summary": clean,
	})
}
