package gchat

import ws "nhooyr.io/websocket"

type Client struct {
	Id int
	Ws *ws.Conn
}
