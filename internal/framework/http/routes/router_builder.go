package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouterBuilder struct {
	StatusRoute  *StatusRoute
	NumbersRoute *NumbersRoute
}

func NewRouterBuilder(
	statusRoute *StatusRoute,
	numbersRoute *NumbersRoute,
) *RouterBuilder {
	return &RouterBuilder{
		StatusRoute:  statusRoute,
		NumbersRoute: numbersRoute,
	}
}

func (r *RouterBuilder) BuildRouter() (*mux.Router, error) {
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

	// Paths for /api/
	apiRouter.HandleFunc(r.StatusRoute.Path, r.StatusRoute.Handle)

	// Paths for /api/v0/
	var apiV0Router = apiRouter.PathPrefix("/v0").Subrouter()
	apiV0Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	var apiV1Router = apiRouter.PathPrefix("/v1").Subrouter()
	apiV1Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	// Paths for /api/v1/
	apiV1Router.HandleFunc(r.NumbersRoute.Path, r.NumbersRoute.CreateNumberController.Execute).Methods("POST")
	apiV1Router.HandleFunc(r.NumbersRoute.PathID, r.NumbersRoute.UpdateNumberByIdController.Execute).Methods("PUT")
	apiV1Router.HandleFunc(r.NumbersRoute.PathID, r.NumbersRoute.RemoveNumberByIdController.Execute).Methods("DELETE")
	apiV1Router.HandleFunc(r.NumbersRoute.PathID, r.NumbersRoute.GetNumberByIdController.Execute)
	apiV1Router.HandleFunc(r.NumbersRoute.Path, r.NumbersRoute.GetAllNumbersController.Execute)

	return router, nil
}
