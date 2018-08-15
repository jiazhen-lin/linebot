package model

import "database/sql"

// User represents the user record
type User struct {
	UserID string
	Type   string
	Name   sql.NullString
}
