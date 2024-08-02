package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("name is required")
var errProjectIDRequired = errors.New("project_id is required")

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

	var task *Task
	err = json.Unmarshal(body, &task) // unmarshal the body of the request into task
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Error unmarshalling task payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Error validating task payload"})
		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}

	WriteJSON(w, http.StatusCreated, t)

}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	return nil

}

func (s *TaskService) GetTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TaskService) GetTasks(w http.ResponseWriter, r *http.Request) {

}

func (s *TaskService) UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TaskService) DeleteTask(w http.ResponseWriter, r *http.Request) {

}
