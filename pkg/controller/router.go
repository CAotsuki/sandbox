package controller

import (
	"net/http"
)

type Router interface {
	HandleTodoRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TodoController
}

func NewRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ro.tc.GetTodos(w, r)
	case http.MethodPost:
		ro.tc.PostTodo(w, r)
	case http.MethodPut:
		ro.tc.PutTodo(w, r)
	case http.MethodDelete:
		ro.tc.DeleteTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}
