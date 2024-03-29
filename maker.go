package token

import "time"

type Maker interface {
	// Creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// Checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)

	// Get Payload data from token
	GetPayload(token string) (*Payload, error)
}
