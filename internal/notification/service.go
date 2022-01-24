package notification

import (
	"context"
	"fmt"

	"github.com/autobrr/autobrr/internal/domain"

	"github.com/containrrr/shoutrrr"
	t "github.com/containrrr/shoutrrr/pkg/types"
)

type Service interface {
	Find(ctx context.Context) ([]domain.Notification, error)
	FindByID(ctx context.Context, id int) (*domain.Notification, error)
	Store(ctx context.Context, n *domain.Notification) (*domain.Notification, error)
	Delete(ctx context.Context, id int) error
	Send(event domain.NotificationEvent, msg string) error
}

type service struct {
	repo domain.NotificationRepo
}

func NewService(repo domain.NotificationRepo) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Find(ctx context.Context) ([]domain.Notification, error) {
	notifications, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (s *service) FindByID(ctx context.Context, id int) (*domain.Notification, error) {
	return nil, nil
}

func (s *service) Store(ctx context.Context, n *domain.Notification) (*domain.Notification, error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return nil
}

// Send a
func (s *service) Send(event domain.NotificationEvent, msg string) error {
	// find notifications for type X

	notifications, err := s.repo.List(context.Background())
	if err != nil {
		return err
	}

	var urls []string

	for _, n := range notifications {
		if !n.Enabled {
			continue
		}

		switch n.Type {
		case domain.NotificationTypeDiscord:
			urls = append(urls, fmt.Sprintf("discord://%v@%v", n.Token, n.Webhook))
		default:
			return nil
		}
	}

	if len(urls) == 0 {
		return nil
	}

	sender, err := shoutrrr.CreateSender(urls...)
	if err != nil {
		return err
	}

	p := t.Params{"title": "TEST"}

	sender.Send(msg, &p)

	return nil
}
