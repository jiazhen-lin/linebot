package model

import (
	"database/sql"
	"time"
)

// Accounting represents the account records
type Accounting struct {
	ID             uint           `db:"id"`
	UserID         string         `db:"user_id"`
	Type           string         `db:"type"`
	CreatedTime    time.Time      `db:"created_time"`
	AccountingTime time.Time      `db:"accounting_time"`
	Price          float64        `db:"price"`
	Purpose        sql.NullString `db:"purpose"`
}
