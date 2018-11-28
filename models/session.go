package models

import (
	"time"
)

func (s *Session) IsAlive() bool {
	return s.TTL > time.Now().Unix()
}