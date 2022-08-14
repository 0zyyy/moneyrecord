package history

type NewHistoryInput struct {
	IDHistory int    `json:"id_history"`
	IDUser    int    `json:"id_user"`
	Type      string `json:"type"`
	Date      string `json:"date"`
	Total     string `json:"total"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Details   string `json:"details"`
}

type Search struct {
	UserID int
	Date   string
}

type Income struct {
	HistorySearch Search `json:"search"`
	Type          string `json:"type"`
}

type DeleteHistory struct {
	IDHistory int
}
