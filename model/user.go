package model

import "database/sql"

// User represents the user record
type User struct {
	UserID   string         `db:"userID"`
	UserType string         `db:"userType"`
	Name     sql.NullString `db:"name"`
}
