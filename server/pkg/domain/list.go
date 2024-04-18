package domain

type ListEntry struct {
	Yandexoid
	Friends int    `json:"friends" db:"friends"`
	Status  string `db:"status" json:"status"`
}
