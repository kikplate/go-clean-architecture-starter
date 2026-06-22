package httpdelivery

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DBPinger interface {
	Ping(ctx context.Context) error
}

type RouterDeps struct {
	Logger         *slog.Logger
	RequestTimeout time.Duration
	DB             DBPinger
	UserHandlers   UserHandlers
}

func NewRouter(d RouterDeps) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(d.RequestTimeout))
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	r.Get("/readyz", func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
		defer cancel()
		if err := d.DB.Ping(ctx); err != nil {
			d.Logger.Error("readyz_failed", slog.Any("err", err))
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("not_ready"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})
	r.Route("/v1", func(r chi.Router) {
		r.Post("/users", d.UserHandlers.PostUser)
		r.Get("/users/{id}", d.UserHandlers.GetUser)
	})
	return r
}
