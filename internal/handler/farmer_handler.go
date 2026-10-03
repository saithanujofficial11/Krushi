package handler

import (
	"encoding/json"
	"krushi-server/internal/models"
	"krushi-server/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type FarmerHandler struct {
	repo repository.FarmerRepository
}

func NewFarmerHandler(repo repository.FarmerRepository) *FarmerHandler {
	return &FarmerHandler{repo: repo}
}

func (h *FarmerHandler) CreateFarmer(w http.ResponseWriter, r *http.Request) {
	var farmer models.Farmer
	if err := json.NewDecoder(r.Body).Decode(&farmer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if farmer.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if err := h.repo.Save(&farmer); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(farmer)
}

func (h *FarmerHandler) GetFarmer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "farmerId")
	if id == "" {
		http.Error(w, "Farmer ID is required", http.StatusBadRequest)
		return
	}

	farmer, err := h.repo.GetByID(id)
	if err != nil {
		if err == repository.ErrFarmerNotFound {
			http.Error(w, "Farmer not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(farmer)
}

func (h *FarmerHandler) ListFarmers(w http.ResponseWriter, r *http.Request) {
	farmers, err := h.repo.ListAll()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(farmers)
}

func (h *FarmerHandler) SearchFarmerByPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		http.Error(w, "Phone number is required", http.StatusBadRequest)
		return
	}

	farmer, err := h.repo.GetByPhone(phone)
	if err != nil {
		if err == repository.ErrFarmerNotFound {
			http.Error(w, "Farmer not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(farmer)
}
