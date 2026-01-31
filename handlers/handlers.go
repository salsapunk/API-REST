package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/salsapunk/API-REST/models"
	"github.com/salsapunk/API-REST/repository"
)

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) TaskHandler {
	return TaskHandler{
		Repo: repo,
	}
}

func (th *TaskHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.Repo.ListAll()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(tasks)
}

func (th *TaskHandler) ListByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/GET/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	task, err := th.Repo.ListByID(id)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(task)
}

func (th *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if _, err := th.Repo.Create(&task); err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (th *TaskHandler) Edit(w http.ResponseWriter, r *http.Request) {
}
