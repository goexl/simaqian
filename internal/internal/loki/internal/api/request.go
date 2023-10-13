package api

type Request struct {
	Streams []*Stream `json:"streams,omitempty"`
}
