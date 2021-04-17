package servmsg

const (
	TWelcome = `[Server] Welcome`
)

type WelcomePayload struct {
	Id string `json:"id"`
}

func Welcome(payload WelcomePayload) Message {
	return Message{
		Kind:    TWelcome,
		Payload: payload,
	}
}
