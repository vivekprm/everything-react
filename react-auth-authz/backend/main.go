package main

import (
	"backend/controllers"
	"backend/templates"
	"backend/views"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	tpl := views.Must(views.ParseFS(templates.FS, "login.gohtml", "tailwind.gohtml"))
	r.Get("/login", controllers.StaticHandler(tpl, nil))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	authC := controllers.AuthRequest{}
	authC.Templates.New = views.Must(views.ParseFS(templates.FS, "login.gohtml", "tailwind.gohtml"))
	r.Post("/login", authC.Login)

	fmt.Println("Starting server on :8000...")
	http.ListenAndServe(":8000", r)
}
