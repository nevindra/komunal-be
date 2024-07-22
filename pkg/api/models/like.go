package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	LikeID     uuid.UUID `json:"like_id"`
	UserID     uuid.UUID `json:"user_id"`
	QuestionID string    `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
}
