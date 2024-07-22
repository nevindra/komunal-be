package models

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	QuestionID  string    `json:"question_id"`
	UserID      uuid.UUID `json:"user_id"`
	PosterID    string    `json:"poster_id"`
	Content     string    `json:"content"`
	IsAnonymous bool      `json:"is_anonymous"`
	CreatedAt   time.Time `json:"created_at"`
}
