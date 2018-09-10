package model

import (
	"database/sql"
	"time"
)

// Accounting represents the account records
type Accounting struct {
	ID             uint           `db:"ID"`
	UserID         string         `db:"userID"`
	Kind           string         `db:"kind"`
	CreatedTime    time.Time      `db:"createdTime"`
	AccountingTime time.Time      `db:"accountingTime"`
	Price          float64        `db:"price"`
	Purpose        sql.NullString `db:"purpose"`
}
