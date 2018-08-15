package model

import (
	"database/sql"
	"time"
)

// Accounting represents the account records
type Accounting struct {
	ID             uint
	UserID         string
	Type           string
	CreatedTime    time.Time
	AccountingTime time.Time
	Price          float64
	Purpose        sql.NullString
}
