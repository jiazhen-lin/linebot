package model

import "database/sql"

// User represents the user record
type User struct {
	UserID string         `db:"user_id"`
	Type   string         `db:"type"`
	Name   sql.NullString `db:"name"`
}
