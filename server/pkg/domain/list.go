package domain

type ListEntry struct {
	Yandexoid
	Status string `db:"status" json:"status"`
}
