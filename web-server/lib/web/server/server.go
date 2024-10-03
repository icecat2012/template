package server

import (
	"context"
	"fmt"
	"net/http"
	"stock/lib/utils"
	"time"
)

type I interface {
	Run() error
	Stop() error
}

type Option struct {
	Port int
}
type Handler struct {
	server    *http.Server
	isRunning bool
}

func New(opt Option) I {
	mux := route()
	return &Handler{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%v", opt.Port),
			Handler: mux,
		},
		isRunning: false,
	}
}

func (p *Handler) Run() error {
	if p.isRunning {
		return utils.ErrServerRunning
	}
	p.isRunning = true
	fmt.Println("Starting server")
	return p.server.ListenAndServe()
}

func (p *Handler) Stop() error {
	if !p.isRunning {
		return utils.ErrServerStopping
	}
	p.isRunning = false
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.server.Shutdown(ctx)
}
