package servmsg

type Message struct {
	Kind    string      `json:"kind"`
	Payload interface{} `json:"payload"`
}
