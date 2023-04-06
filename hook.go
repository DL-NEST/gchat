package gchat

// Hook 钩子
type Hook interface {
	OnMessageSend()
	OnMessageStorage()
}

type hook struct {
}

var defaultHook hook

func (h hook) OnMessageSend() {

}

func (h hook) OnMessageStorage() {

}
