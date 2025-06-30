package task

import (
	"encoding/json"
	"io"
	models_task "microservice/Models/task"
	"net/http"
	"strconv"
)

//type service interface {
//	Insertask(t models_task.Task) (string, error)
//	Getalltask() ([]models_task.Task, error)
//	Gettaskbyid(id int) (*models_task.Task, error)
//	Deletetask(id int) (string, error)
//	Completetask(id int) (string, error)
//}

type handler struct {
	service service
}

func New(s service) *handler {
	return &handler{service: s}
}

// Addtask handles creating a new task
// @Summary Add a new task
// @Description Add a new task to the task list
// @Accept json
// @Produce json
// @Param task body models_task.Task true "Task Info"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /task [post]

func (h *handler) Addtask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	var t models_task.Task
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "invalid JSON format", http.StatusBadRequest)
		return
	}

	msg, err := h.service.Insertask(t)

	if err != nil {
		http.Error(w, msg+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": msg,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "failed to generate JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

// Gettaskbyid returns a task by ID
// @Summary Get a task by ID
// @Description Fetch a task by its ID
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models_task.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /task/{id} [get]

func (h *handler) Gettaskbyid(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "error to parsing id", http.StatusInternalServerError)
		return
	}
	var t *models_task.Task
	t, err = h.service.Gettaskbyid(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	response, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "error generating JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

// Getalltask returns all tasks
// @Summary Get all tasks
// @Description Fetch all tasks from the system
// @Produce json
// @Success 200 {array} models_task.Task
// @Failure 500 {object} map[string]string
// @Router /task [get]

func (h *handler) Getalltask(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.Getalltask()
	if err != nil {
		http.Error(w, "failed to fetch tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "failed to marshal tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

// Deletetask deletes a task by ID
// @Summary Delete a task
// @Description Delete a task by its ID
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /task/{id} [delete]

func (h *handler) Deletetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}
	msg, err := h.service.Deletetask(id)
	if err != nil {
		http.Error(w, "failed to delete task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": msg}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

// Completetask marks a task as complete
// @Summary Complete a task
// @Description Mark a task as completed by its ID
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /task/{id} [patch]

func (h *handler) Completetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}
	msg, err := h.service.Completetask(id)
	if err != nil {
		http.Error(w, "failed to complete task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": msg}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}
