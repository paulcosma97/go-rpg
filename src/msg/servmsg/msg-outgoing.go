package servmsg

const (
	TWelcome = `[Server] Welcome`
)

type WelcomePayload struct {
	A string `json:"a"`
}

func Welcome(payload WelcomePayload) Message {
	return Message{
		Kind:    TWelcome,
		Payload: payload,
	}
}
