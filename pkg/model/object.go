package model

import (
	"github.com/google/uuid"
	"time"
)

type Object struct {
	ID      uuid.UUID
	Name    string
	Value   float64
	Created time.Time
}

//NewObject return Object with predefined ID and Created time
func NewObject() Object {
	return Object{
		ID:      uuid.New(),
		Created: time.Now(),
	}
}
