package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`         // each payload has each different ID
	Username  string    `json:"username"`   // Username to know which is used to identify the token owner
	IssueAt   time.Time `json:"issue_at"`   // Time create this token
	ExpiredAt time.Time `json:"expired_at"` // Time this token expired
}

// NewPayload creates new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error){
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID: tokenID,
		Username: username,
		IssueAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}