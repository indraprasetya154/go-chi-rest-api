package todo

import (
	"github.com/go-chi/chi/v5"
	"github.com/indraprasetya154/go-chi-rest-api/src/modules/todo/handlers"
)

func TodoRoutes() chi.Router {
	r := chi.NewRouter()
	todoHandler := handlers.TodoHandler{}
	r.Get("/", todoHandler.ListTodos)
	// r.Post("/", todoHandler.CreateTodos)
	// r.Get("/{id}", todoHandler.GetTodos)
	// r.Put("/{id}", todoHandler.UpdateTodos)
	// r.Delete("/{id}", todoHandler.DeleteTodos)
	return r
}
