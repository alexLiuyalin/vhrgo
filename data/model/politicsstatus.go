package model

import "database/sql"

type PoliticsStatus struct {
	ID   int64          `json:"id" db:"id"`
	Name sql.NullString `json:"name" db:"name"`
}
