package server

import (
	"fmt"
	"log"

	"net/http"

	"github.com/eduardogr/webser-go/pkg/config"
)

type Server struct {
	router *Router
}

func NewServer(router *Router) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) HandleRequests() error {
	router, err := s.router.BuildRouter()

	if err != nil {
		log.Fatal(err)
		return err
	}

	c := config.GetConfiguration(config.STRATEGY)
	apiPort := fmt.Sprintf(":%d", c.ApiPort)
	err = http.ListenAndServe(apiPort, router)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
