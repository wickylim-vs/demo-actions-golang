package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ts *template.Template

func init() {
	ts = template.Must(template.ParseFiles("./ui/html/home.tmpl"))
}

func main() {

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(measureResponseDuration)

	// Routes
	r.Get("/", home)
	r.Get("/ping", ping)

	filesDir := http.Dir("./ui/static/")
	fileServer(r, "/static", filesDir)

	r.Handle("/metrics", promhttp.Handler())

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	opsRequest.Inc()
	hostname, _ := os.Hostname()
	ts.Execute(w, hostname)
}

func ping(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	opsRequest.Inc()
	fmt.Fprintf(w, "pong from %v", hostname)
}
