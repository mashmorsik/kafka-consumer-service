package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mashmorsik/kafka-consumer-service/config"
	mw "github.com/mashmorsik/kafka-consumer-service/pkg/middleware"
	"github.com/mashmorsik/logger"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

type HTTPServer struct {
	Config *config.Config
}

func NewHTTPServer(config *config.Config) *HTTPServer {
	return &HTTPServer{Config: config}
}

func (s *HTTPServer) StartServer(ctx context.Context) error {
	r := chi.NewRouter()
	r.Use(mw.LoggingMiddleware)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Post("/kafka", s.SendMessageToKafka)

	logger.Infof("HTTPServer is listening on port: %s\n", s.Config.Server.Port)

	httpServer := &http.Server{
		Addr:              s.Config.Server.Port,
		Handler:           cors.AllowAll().Handler(r),
		ReadHeaderTimeout: 30 * time.Second,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(ctx)
	})

	if err := g.Wait(); err != nil {
		return errors.WithMessagef(err, "exit reason: %s \n", err)
	}

	return nil
}

func (s *HTTPServer) SendMessageToKafka(w http.ResponseWriter, r *http.Request) {

}
