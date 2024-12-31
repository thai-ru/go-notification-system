package notification

import "time"

type Notification struct {
	ID        uint  `gorm:"primary_key"`
	UserID    int64 `gorm:"index"`
	Title     string
	Content   string
	IsRead    bool
	CreatedAt time.Time
}
