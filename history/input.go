package history

type NewHistoryInput struct {
	IDHistory int    `json:"id_history"`
	IDUser    int    `json:"id_user" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Total     string `json:"total" binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Details   string `json:"details" binding:"required"`
}

type Search struct {
	IDUser int    `json:"id_user" binding:"required"`
	Date   string `json:"date"`
}

type Income struct {
	HistorySearch Search `json:"search" binding:"required"`
	Type          string `json:"type" binding:"required"`
}

type DeleteHistory struct {
	IDHistory int
}
