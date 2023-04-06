package gchat

import ws "nhooyr.io/websocket"

type User struct {
	Id         uint32    `json:"id"`
	ClientPool []*Client `json:"clientPool"`
}

func (c *User) AddClient(ws *ws.Conn) {
	c.ClientPool = append(c.ClientPool, &Client{
		Id: 0,
		Ws: ws,
	})
}

// AddSession 添加会话
// SessionType group | user
func (c *User) AddSession(ws *ws.Conn) {
	c.ClientPool = append(c.ClientPool, &Client{
		Id: 0,
		Ws: ws,
	})
}

// SendMsg 发送消息
func (c *User) SendMsg(ws *ws.Conn) {
	c.ClientPool = append(c.ClientPool, &Client{
		Id: 0,
		Ws: ws,
	})
}
