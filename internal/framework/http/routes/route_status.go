package routes

import (
	"log"
	"net/http"
)

type StatusRoute struct {
	Path string
}

func NewStatusRoute() *StatusRoute {
	return &StatusRoute{
		Path: "/status",
	}
}

func (route *StatusRoute) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("HTTP/1.1 200 OK")
}
