package models

import (
	"time"

	"github.com/google/uuid"
)

type Share struct {
	ShareID    uuid.UUID `json:"share_id"`
	UserID     uuid.UUID `json:"user_id"`
	QuestionID string    `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
}
