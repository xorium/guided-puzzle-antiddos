package tcp

type MsgType byte

const (
	MsgTypeRegister MsgType = iota
	MsgTypeTourFailed
	MsgTypeVerified
	MsgTypeHeartBeat
	MsgTypeBroadcastGuides
)

type GeneralRequestMsg struct{
	Type MsgType
}

type RegisteredRespMsg struct {
	ID int
	SharedKey string
}

type GuideInfo struct {
	SharedKey string
	PublicAddr string
}

type BroadcastGuidesMsg struct {
	ActiveGuides []GuideInfo
}

func getEmptyMessageByType(tp MsgType) any {
	switch tp {
	case MsgTypeRegister:
		return &
	}
}