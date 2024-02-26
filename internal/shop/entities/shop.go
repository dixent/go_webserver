package entities

import "time"

type Shop struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name" json:"name"`
	OwnerId   int64     `db:"owner_id" json:"owner_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
