package api

type Stream struct {
	Stream map[string]string `json:"stream,omitempty"`
	Values []*Value          `json:"values,omitempty"`
}
