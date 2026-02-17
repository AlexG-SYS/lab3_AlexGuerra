package routes

import (
	"net/http"

	"github.com/AlexG-SYS/lab3_AlexGuerra/internal/handlers"
	"github.com/AlexG-SYS/lab3_AlexGuerra/internal/helpers"
)

func SetupRoutes() http.Handler {
	app := &helpers.Application{}
	h := &handlers.Handler{App: app}

	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HomeHandler)
	mux.HandleFunc("/home", h.HomeHandler)
	mux.HandleFunc("/login", h.LoginHandler)

	return mux
}
