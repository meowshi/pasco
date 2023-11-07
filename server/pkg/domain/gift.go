package domain

type GetRes struct {
	LuckyLogin     string      `json:"lucky_login"`
	Himself        bool        `json:"himself"`
	CollectorLogin string      `json:"collector_login"`
	CollectedCount int         `json:"collected_count"`
	FilterUrl      string      `json:"filter_url"`
	CollectedFor   []string    `json:"collected_for"`
	CollectorFor   interface{} `json:"collector_for"`
	Key            string      `json:"key"`
}
