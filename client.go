package gchat

import ws "nhooyr.io/websocket"

type Client struct {
	Id uint32   `json:"id"`
	Ws *ws.Conn `json:"-"`
}
