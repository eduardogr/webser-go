package http

import (
	"fmt"
	"log"

	"net/http"

	"github.com/eduardogr/webser-go/internal/application/config"
	"github.com/eduardogr/webser-go/internal/framework/http/routes"
)

type Server struct {
	RouterBuilder *routes.RouterBuilder
}

func NewServer(routerBuilder *routes.RouterBuilder) *Server {
	return &Server{
		RouterBuilder: routerBuilder,
	}
}

func (s *Server) HandleRequests() error {
	router, err := s.RouterBuilder.BuildRouter()

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
