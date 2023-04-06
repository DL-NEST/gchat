package main

import "gchat"

type hook struct {
}

func (h hook) OnMessageSend() {
	//TODO implement me
	panic("implement me")
}

func (h hook) OnMessageStorage() {
	//TODO implement me
	panic("implement me")
}

func main() {
	_ = gchat.NewHub(
		gchat.WithHook(hook{}))
}
