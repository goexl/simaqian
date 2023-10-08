package loki

type Stream struct {
	Labels map[string]string `json:"stream"`
	Values [][2]string       `json:"values"`
}
