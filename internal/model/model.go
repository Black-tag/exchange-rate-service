package model






type RateResponse struct {
	Success   bool               `json:"success"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp int64              `json:"timestamp"`
	Source    string             `json:"source"`
	Quotes    map[string]float64 `json:"quotes"`

}

type ConvertResponse struct {
	Success bool   `json:"success"`
    Terms   string `json:"terms"`
    Privacy string `json:"privacy"`

    Query struct {
        From   string  `json:"from"`
        To     string  `json:"to"`
        Amount float64 `json:"amount"`
    } `json:"query"`

    Info struct {
        Timestamp int64   `json:"timestamp"`
        Quote     float64 `json:"quote"`
    } `json:"info"`

    Result float64 `json:"result"`

}