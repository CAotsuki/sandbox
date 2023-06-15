package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"sandbox/pkg/controller/dto"
	"sandbox/pkg/model/entity"
	"sandbox/pkg/model/repository"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request, db *sql.DB)
	PostTodo(w http.ResponseWriter, r *http.Request, db *sql.DB)
	PutTodo(w http.ResponseWriter, r *http.Request, db *sql.DB)
	DeleteTodo(w http.ResponseWriter, r *http.Request, db *sql.DB)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	todos, err := tc.tr.GetTodos(db)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoResponses []dto.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{
			Id:      todo.Id,
			Title:   todo.Title,
			Content: todo.Content,
		})
	}

	var todosResponse dto.TodosResponse
	todosResponse.Todos = todoResponses

	output, _ := json.MarshalIndent(todosResponse.Todos, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	json.Unmarshal(body, &todoRequest)

	todo := entity.TodoEntity{
		Title:   todoRequest.Title,
		Content: todoRequest.Content,
	}

	id, err := tc.tr.InsertTodo(todo, db)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (tc *todoController) PutTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	json.Unmarshal(body, &todoRequest)

	todo := entity.TodoEntity{
		Id:      todoId,
		Title:   todoRequest.Title,
		Content: todoRequest.Content,
	}

	err = tc.tr.UpdateTodo(todo, db)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
	}

	err = tc.tr.DeleteTodo(todoId, db)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
