package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskService struct {
	store Store
}

func NewTaskService(store Store) *TaskService {
	return &TaskService{
		store: store,
	}
}

func (s *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.GetTask).Methods("GET")
	r.HandleFunc("/tasks", s.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", s.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", s.DeleteTask).Methods("DELETE")

}

func (s *TaskService) CreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body) // read the body of the request and store it in body and err
	fmt.Println(string(body))       // print the body of the request
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close() // close the body of the request

}

func (s *TaskService) GetTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TaskService) GetTasks(w http.ResponseWriter, r *http.Request) {

}

func (s *TaskService) UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TaskService) DeleteTask(w http.ResponseWriter, r *http.Request) {

}
