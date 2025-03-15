package e11

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/vladopajic/go-actor/actor"
)

var _ actor.Actor = (*service)(nil) // ensure that service implements Actor

func NewHTTPService(
	addr string,
	handler http.Handler,
) *service {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &service{httpServer}
}

type service struct {
	httpServer *http.Server
}

func (s *service) Start() {
	log.Info().Msg("starting HTTP service")

	// Listen and serve in another goroutine
	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("failed starting http server")
		}
	}()
}

func (s *service) Stop() {
	log.Info().Msg("stopping HTTP service")

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed shutting down http server")
	}
}
