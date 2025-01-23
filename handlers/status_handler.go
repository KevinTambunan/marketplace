package handlers

import (
	"encoding/json"
	"marketplace/models"
	"marketplace/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StatusHandler struct {
	Service *services.StatusService
}

func NewStatusHandler(service *services.StatusService) *StatusHandler {
	return &StatusHandler{Service: service}
}

func (h *StatusHandler) GetAllStatuses(w http.ResponseWriter, r *http.Request) {
	statuses, err := h.Service.GetAllStatuses()
	if err != nil {
		http.Error(w, "Failed to retrieve statuses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func (h *StatusHandler) CreateStatus(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	statusID, err := h.Service.CreateStatus(request.Name)
	if err != nil {
		http.Error(w, "Failed to create status", http.StatusInternalServerError)
		return
	}

	newStatus, err := h.Service.GetStatusById(statusID)
	if err != nil {
		http.Error(w, "Failed to retrieve new status", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string         `json:"message"`
		Status  *models.Status `json:"status"`
	}{
		Message: "Status created successfully",
		Status:  newStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *StatusHandler) GetStatusById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid status ID", http.StatusBadRequest)
		return
	}

	status, err := h.Service.GetStatusById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (h *StatusHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid status ID", http.StatusBadRequest)
		return
	}

	var request struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateStatus(id, request.Name); err != nil {
		http.Error(w, "Failed to update status", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Status updated successfully"))
}

func (h *StatusHandler) DeleteStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid status ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteStatus(id); err != nil {
		http.Error(w, "Failed to delete status", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Status deleted successfully"))
}
