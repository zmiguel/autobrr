package database

import (
	"context"
	"database/sql"
	"github.com/autobrr/autobrr/internal/domain"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type NotificationRepo struct {
	db *SqliteDB
}

func NewNotificationRepo(db *SqliteDB) domain.NotificationRepo {
	return &NotificationRepo{
		db: db,
	}
}

func (r *NotificationRepo) List(ctx context.Context) ([]domain.Notification, error) {

	rows, err := r.db.handler.QueryContext(ctx, "SELECT id, name, type, enabled, events, token, api_key, webhook, title, icon, host, username, password, channel, targets, devices, created_at, updated_at FROM notification ORDER BY name ASC")
	if err != nil {
		log.Error().Stack().Err(err).Msg("filters_list: error query data")
		return nil, err
	}

	defer rows.Close()

	var notifications []domain.Notification
	for rows.Next() {
		var n domain.Notification

		var token, apiKey, webhook, title, icon, host, username, password, channel, targets, devices sql.NullString
		if err := rows.Scan(&n.ID, &n.Name, &n.Type, &n.Enabled, pq.Array(&n.Events), &token, &apiKey, &webhook, &title, &icon, &host, &username, &password, &channel, &targets, &devices, &n.CreatedAt, &n.UpdatedAt); err != nil {
			log.Error().Stack().Err(err).Msg("notification_list: error scanning data to struct")
			return nil, err
		}

		n.Token = token.String
		n.APIKey = apiKey.String
		n.Webhook = webhook.String
		n.Title = title.String
		n.Icon = icon.String
		n.Host = host.String
		n.Username = username.String
		n.Password = password.String
		n.Channel = channel.String
		n.Targets = targets.String
		n.Devices = devices.String

		notifications = append(notifications, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *NotificationRepo) FindByID(ctx context.Context, id int) (*domain.Notification, error) {

	row := r.db.handler.QueryRowContext(ctx, "SELECT id, name, type, enabled, events, token, api_key, webhook, title, icon, host, username, password, channel, targets, devices, created_at, updated_at FROM notification WHERE id = ?", id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var n domain.Notification

	var token, apiKey, webhook, title, icon, host, username, password, channel, targets, devices sql.NullString
	if err := row.Scan(&n.ID, &n.Name, &n.Type, &n.Enabled, pq.Array(&n.Events), &token, &apiKey, &webhook, &title, &icon, &host, &username, &password, &channel, &targets, &devices, &n.CreatedAt, &n.UpdatedAt); err != nil {
		log.Error().Stack().Err(err).Msgf("notification: %v : error scanning data to struct", id)
		return nil, err
	}

	n.Token = token.String
	n.APIKey = apiKey.String
	n.Webhook = webhook.String
	n.Title = title.String
	n.Icon = icon.String
	n.Host = host.String
	n.Username = username.String
	n.Password = password.String
	n.Channel = channel.String
	n.Targets = targets.String
	n.Devices = devices.String

	return &n, nil
}

func (r *NotificationRepo) Store(ctx context.Context, notification domain.Notification) (*domain.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (r *NotificationRepo) Delete(ctx context.Context, notificationID int) error {
	_, err := r.db.handler.ExecContext(ctx, `DELETE FROM notification WHERE id = ?`, notificationID)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error executing query")
		return err
	}

	log.Info().Msgf("notification.delete: successfully deleted: %v", notificationID)

	return nil
}
