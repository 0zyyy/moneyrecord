package history

type ResponseHistory struct {
	IDHistory int    `json:"id_history"`
	Date      string `json:"date"`
	Total     string `json:"total"`
	Type      string `json:"type"`
	Details   string `json:"details"`
}

type ResponseAnalysis struct {
	Today     float64
	Yesterday float64
	Week      []float64
	Month     MonthResult
}

type MonthResult struct {
	Income  float64
	Outcome float64
}

func ResponseHistoryFormatter(history History) ResponseHistory {
	response := ResponseHistory{
		IDHistory: history.IDHistory,
		Date:      history.Date,
		Total:     history.Total,
		Type:      history.Type,
		Details:   history.Details,
	}
	return response
}
