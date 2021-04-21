package cmsg

const (
	// Not actually sent by the client, but simulated
	TConnect    = `[Client] Connect`
	TPing       = `[Client] Ping`
	TSetProfile = `[Client] Set Profile`
	TCreateMatch = `[Client] Create Match`
	TJoinMatch = `[Client] Join Match`
	TLeaveMatch = `[Client] Leave Match`

)

type EmptyMsg struct{}

type SetProfilePayload struct {
	DisplayName string `json:"displayName"`
}

type JoinMatchPayload = string
type LeaveMatchPayload = string