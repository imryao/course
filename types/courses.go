package types

import "time"

type Course struct {
	ID         *int64     `json:"id"`
	Name       *string    `json:"name"`
	Quota      *int64     `json:"quota"`
	StartedAt  *time.Time `json:"started_at"`
	EndedAt    *time.Time `json:"ended_at"`
	InsertedAt *time.Time `json:"inserted_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
