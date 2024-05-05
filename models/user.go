package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	ScreenName string    `json:"screenName"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}
