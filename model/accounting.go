package model

import (
	"database/sql"
	"time"
)

// Accounting represents the account records
type Accounting struct {
	ID             uint           `db:"ID"`
	UserID         string         `db:"userID"`
	Category       string         `db:"category"`
	CreatedTime    time.Time      `db:"createdTime"`
	AccountingTime time.Time      `db:"accountingTime"`
	Price          float64        `db:"price"`
	Detail         sql.NullString `db:"detail"`
}
