package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aserold/go-crud/config"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

const (
	ctxTimeout = 5
)

type Server struct {
	echo *echo.Echo
	cfg  *config.Config
	db   *sqlx.DB
}

func NewServer(cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         s.cfg.Server.Port,
		ReadTimeout:  time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * s.cfg.Server.WriteTimeout,
	}

	go func() {
		log.Printf("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			log.Fatalf("error starting Server: %s", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	log.Println("server exited properly")
	return s.echo.Server.Shutdown(ctx)
}
