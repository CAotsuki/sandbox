package main

import (
	"fmt"
	"net/http"

	"sandbox/pkg/controller"
	"sandbox/pkg/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", homeLink)
	http.HandleFunc("/todos/", ro.HandleTodoRequest)
	server.ListenAndServe()
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
