package notification

import (
	"context"
	"fmt"
	"log"
	"notification-system/internal/email"
	"strconv"
	"time"
)

type Service interface {
	CreateNotification(ctx context.Context, userID int64, title, content string, sendEmail bool) (*Notification, error)
	GetUserNotifications(ctx context.Context, userID int64) ([]Notification, error)
}

type service struct {
	repo         Repository
	emailService email.Service
}

func NewService(repo Repository, emailService email.Service) Service {
	return &service{
		repo:         repo,
		emailService: emailService,
	}
}

func (s *service) CreateNotification(ctx context.Context, userID int64, title, content string, sendEmail bool) (*Notification, error) {
	notification := &Notification{
		UserID:    userID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(ctx, notification); err != nil {
		return nil, err
	}

	if sendEmail {
		// In a real life scenario get email from user service
		userEmail := fmt.Sprintf("thairurobin@gmail.com", userID)
		go func() {
			err := s.emailService.SendEmail(userEmail, title, content)
			if err != nil {
				log.Println(err)
			}
		}()
	}
	return notification, nil
}

func (s *service) GetUserNotifications(ctx context.Context, userID int64) ([]Notification, error) {
	return s.repo.GetUserNotifications(ctx, strconv.FormatInt(userID, 10))
}
