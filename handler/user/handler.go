package user

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	models_user "microservice/Models/user"
)

type service interface {
	InsertUser(u models_user.User) (string, error)
	GetUserByID(id int) (*models_user.User, error)
	GetAllUsers() ([]models_user.User, error)
	DeleteAllUsers() (string, error)
	DeleteUserByID(id int) (string, error)
}

type handler struct {
	service service
}

func New(s service) *handler {
	return &handler{service: s}
}

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	var u models_user.User
	err = json.Unmarshal(bodyBytes, &u)
	if err != nil {
		http.Error(w, "Failed to unmarshal request JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := h.service.InsertUser(u)
	if err != nil {
		http.Error(w, msg+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	resBody := map[string]string{"message": msg}
	jsonBytes, err := json.Marshal(resBody)
	if err != nil {
		http.Error(w, "Failed to marshal response JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to get users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to marshal users JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

func (h *handler) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	msg, err := h.service.DeleteAllUsers()
	if err != nil {
		http.Error(w, "Failed to delete users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resBody := map[string]string{"message": msg}
	jsonBytes, err := json.Marshal(resBody)
	if err != nil {
		http.Error(w, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}

func (h *handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	msg, err := h.service.DeleteUserByID(id)
	if err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resBody := map[string]string{"message": msg}
	jsonBytes, err := json.Marshal(resBody)
	if err != nil {
		http.Error(w, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, "Error to sending  responses", http.StatusInternalServerError)
		return
	}
}
