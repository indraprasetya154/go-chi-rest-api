package handlers

import (
	"net/http"

	"github.com/indraprasetya154/go-chi-rest-api/src/modules/todo/services"
)

type TodoHandler struct {
	Service *services.TodoService
}

func (t *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	t.Service.ListTodo(w, r)
}
