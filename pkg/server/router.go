package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{
		handler: handler,
	}
}

func (r *Router) BuildRouter() (*mux.Router, error) {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	var apiRouter = router.PathPrefix("/api").Subrouter()
	apiRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	// middleware to our subrouter to print what route is currently requested.
	apiRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})

	var apiV0Router = apiRouter.PathPrefix("/v0").Subrouter()
	apiV0Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	var apiV1Router = apiRouter.PathPrefix("/v1").Subrouter()
	apiV1Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	// Paths for /api/v1/
	apiV1Router.HandleFunc("/numbers", r.handler.CreateNumber).Methods("POST")
	apiV1Router.HandleFunc("/numbers/{id}", r.handler.UpdateNumber).Methods("PUT")
	apiV1Router.HandleFunc("/numbers/{id}", r.handler.DeleteNumber).Methods("DELETE")
	apiV1Router.HandleFunc("/numbers/{id}", r.handler.GetNumber)
	apiV1Router.HandleFunc("/numbers", r.handler.GetAllNumbers)

	return router, nil
}
