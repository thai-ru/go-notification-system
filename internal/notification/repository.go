package notification

import (
	"context"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, notification *Notification) error
	GetUserNotifications(ctx context.Context, UserID string) ([]Notification, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, notification *Notification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

func (r *repository) GetUserNotifications(ctx context.Context, UserID string) ([]Notification, error) {
	var notifications []Notification
	err := r.db.WithContext(ctx).Where("user_id = ?", UserID).Find(&notifications).Error
	return notifications, err
}
