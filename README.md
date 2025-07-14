
# 🎓 FealtyX - Student Management REST API

FealtyX is a simple RESTful API built in Go for managing student data, including a feature to generate student summaries using an LLM (via Ollama).

---

## 🧰 Tech Stack

- 🐹 Go (Golang)
- 🧰 Gorilla Mux Router
- 🧠 Ollama LLM API (e.g., `gemma:2b`)

---

## 📁 Project Structure

```

fealtyX/
├── handlers/         # API route handlers
├── models/           # Data models and in-memory store
├── main.go           # Entry point of the application
├── go.mod / go.sum   # Go dependencies
└── README.md

````

---

## 🚀 Getting Started

### 1. Install & Run Ollama

> You must have Ollama installed and running locally.

```bash
ollama run gemma:2b
````

This launches the LLM model locally on `http://localhost:11434`.

---

### 2. Run the Go API

```bash
go run main.go
```

By default, the server runs on:
**`http://localhost:8080`**

---

## 📡 API Endpoints

| Method | Endpoint                 | Description                         |
| ------ | ------------------------ | ----------------------------------- |
| POST   | `/students`              | Create a new student                |
| GET    | `/students`              | Get all students                    |
| GET    | `/students/{id}`         | Get a specific student              |
| PUT    | `/students/{id}`         | Update a student                    |
| DELETE | `/students/{id}`         | Delete a student                    |
| GET    | `/students/{id}/summary` | Generate student summary via Ollama |

---

## 📌 Sample Request - Create Student

```bash
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{"name":"John Doe", "age":21, "email":"john@example.com"}'
```

---

## 📌 Sample Request - Get Summary

```bash
curl http://localhost:8080/students/1/summary
```

**Sample Response:**

```json
{
  "summary": "Name: John Doe, Age: 21, Email: john@example.com"
}
```

---

