package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//
// Session models a session inside
// of santata
//
type Session struct {
	ID        uuid.UUID              `json:"id"`
	Token     string                 `json:"token"`
	Data      map[string]interface{} `json:"data"`
	ExpiresAt time.Time              `json:"expires_at"`
}
