package cmsg

const (
	TConnect = `[Client] Connect`
	TPing = `[Client] Ping`
)

type EmptyMsg struct{}

func Ping() *Message {
	return &Message{
		Kind: TPing,
	}
}

func Connect() *Message {
	return &Message{
		Kind: TConnect,
	}
}
