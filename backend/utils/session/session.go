package session

import (
	"time"
)

type Session struct {
	Username string
	Expiry   time.Time
	Role     string
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

var Sessions = map[string]Session{}
