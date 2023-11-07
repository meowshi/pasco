package domain

type LockerEvent struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type PrintBraceletReq struct {
	EventId    string `json:"event_id"`
	PrinterId  string `json:"printer_id"`
	PrintCount int    `json:"print_count"`
}

type Printer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
