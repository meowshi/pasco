package domain

import "time"

type Env struct {
	Key      string    `db:"key" json:"key"`
	Value    string    `db:"value" json:"value"`
	EditedAt time.Time `db:"edited_at" json:"edited_at"`
}
