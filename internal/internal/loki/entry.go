package loki

type Entry struct {
	Level     string  `json:"level"`
	Timestamp float64 `json:"ts"`
	Message   string  `json:"msg"`
	Caller    string  `json:"caller"`

	raw string
}
