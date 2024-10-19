package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/indraprasetya154/go-chi-rest-api/src/modules/todo/models"
)

type TodoService struct{}

func (t *TodoService) ListTodo(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.ListTodos())
	if err != nil {
		log.Printf("Error encoding todos: %v", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
