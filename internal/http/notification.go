package http

import (
	"context"
	"github.com/autobrr/autobrr/internal/domain"
	"github.com/go-chi/chi"
	"net/http"
)

type notificationService interface {
	Find(ctx context.Context) ([]domain.Notification, error)
	FindByID(ctx context.Context, id int) (*domain.Notification, error)
	Store(ctx context.Context, n *domain.Notification) (*domain.Notification, error)
	Delete(ctx context.Context, id int) error
}

type notificationHandler struct {
	encoder encoder
	service notificationService
}

func newNotificationHandler(encoder encoder, service notificationService) *notificationHandler {
	return &notificationHandler{
		encoder: encoder,
		service: service,
	}
}

func (h notificationHandler) Routes(r chi.Router) {
	r.Get("/", h.list)
}

func (h notificationHandler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := h.service.Find(ctx)
	if err != nil {
		h.encoder.StatusNotFound(ctx, w)
		return
	}

	h.encoder.StatusResponse(ctx, w, list, http.StatusOK)
}
