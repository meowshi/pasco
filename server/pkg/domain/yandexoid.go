package domain

type Yandexoid struct {
	Login   string `db:"login" json:"login"`
	Name    string `db:"name" json:"name"`
	Surname string `db:"surname" json:"surname"`
}

type YandexoidRegs struct {
	*Yandexoid
	Key    string                          `json:"key"`
	PickId int64                           `json:"pick_id"`
	Events []*EventWithYandexoidStatusCell `json:"events"`
}

type EventWithYandexoidStatusCell struct {
	*Event
	StatusCell string `db:"status_cell" json:"status_cell"`
}
