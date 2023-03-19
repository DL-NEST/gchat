package gchat

import ws "nhooyr.io/websocket"

type User struct {
	Id         int
	ClientPool map[int]*Client
}

func (c *User) AddClient(clientId int, ws *ws.Conn) {
	c.ClientPool[clientId] = &Client{
		Id: clientId,
		Ws: ws,
	}
}
