package loki

type Request struct {
	Streams []Stream `json:"streams"`
}
