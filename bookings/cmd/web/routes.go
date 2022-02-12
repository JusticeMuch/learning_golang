package main

import (
	"net/http"

	"github.com/JusticeMuch/bookings/pkg/config"
	"github.com/JusticeMuch/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole) test middleware
	mux.Use(NoSurve)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/gudetama", handlers.Repo.Gudetama)
	mux.Get("/gintama", handlers.Repo.Gintama)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
