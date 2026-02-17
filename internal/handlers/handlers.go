package handlers

import (
	"net/http"

	"github.com/AlexG-SYS/lab3_AlexGuerra/internal/helpers"
)

// Use the Application struct from helpers as a receiver
type Handler struct {
	App *helpers.Application
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":       "available",
		"productSKU":   "72936734897923",
		"product_name": "White Shirt",
		"cost":         "23.23",
		"price":        "49.99",
		"qty":          "5",
	}
	h.App.WriteJSON(w, http.StatusOK, data, nil)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Logic example: check if fields are missing
	if input.Email == "" || input.Password == "" {
		h.App.ErrorJSON(w, http.StatusUnprocessableEntity, "missing credentials")
		return
	}

	h.App.WriteJSON(w, http.StatusOK, map[string]string{"message": "login successful"}, nil)
}
