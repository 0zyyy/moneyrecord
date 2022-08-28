package history

type ResponseHistory struct {
	IDHistory int    `json:"id_history"`
	Date      string `json:"date"`
	Total     string `json:"total"`
	Type      string `json:"type"`
	Details   string `json:"details"`
}

type ResponseAnalysis struct {
	Today     float64     `json:"today"`
	Yesterday float64     `json:"yesterday"`
	Week      []float64   `json:"week"`
	Month     MonthResult `json:"month"`
}

type MonthResult struct {
	Income  float64 `json:"income"`
	Outcome float64 `json:"outcome"`
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
