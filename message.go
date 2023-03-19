package gchat

type MessageType int

const (
	Text MessageType = iota
	Voice
	Picture
	Video
)

// Message 消息
type Message struct {
	Type    MessageType
	Payload []byte
}
