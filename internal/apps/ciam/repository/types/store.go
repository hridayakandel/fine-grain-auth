package types

import "time"

type Store struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at_timestamp"`
	UpdatedAt time.Time `db:"updated_at_timestamp"`
	DeletedAt time.Time `db:"deleted_at_timestamp"`
}
