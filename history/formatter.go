package history

type ResponseHistory struct {
	IDHistory int    `json:"id_history"`
	Date      string `json:"date"`
	Total     string `json:"total"`
	Type      string `json:"type"`
}

type ResponseAnalysis struct {
	Today     uint
	Yesterday uint
	Week      uint
	Month     MonthResult
}

type MonthResult struct {
	Income  uint
	Outcome uint
}
