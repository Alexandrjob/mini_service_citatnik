package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"math/rand"
	"mini_service_citatnik/src/internal/models"
	"mini_service_citatnik/src/internal/models/dto"
	"mini_service_citatnik/src/internal/service"
	"mini_service_citatnik/src/pkg/db"
	"net/http"
	"strconv"
	"time"
)

type QuoteHandler struct {
	service *service.QuoteService
}

func NewQuoteHandler(service *service.QuoteService) *QuoteHandler {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &QuoteHandler{service}
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req dto.Quote
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Валидация
	if req.Author == "" || req.Text == "" {
		http.Error(w, "Author and quote text are required", http.StatusBadRequest)
		return
	}

	quote := &models.Quote{
		Author: req.Author,
		Text:   req.Text,
	}

	// Сохранение через репозиторий
	if err := h.service.Insert(r.Context(), quote); err != nil {
		http.Error(w, "Failed to create quote", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func (h *QuoteHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to get quotes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quotes, err := h.service.GetAll(r.Context())
	if err != nil || len(quotes) == 0 {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}

	randomQuote := quotes[rand.Intn(len(quotes))]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randomQuote)
}

func (h *QuoteHandler) GetByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author == "" {
		http.Error(w, "Author parameter is required", http.StatusBadRequest)
		return
	}

	quotes, err := h.service.GetByAuthor(r.Context(), author)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Quotes not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to get quotes", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid quote ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Quote not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete quote", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
