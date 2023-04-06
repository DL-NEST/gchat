package gchat

import (
	ws "nhooyr.io/websocket"
)

type Hub struct {
	// 是否持久化
	Persistence bool
	// 持久化方式
	PersistenceType int
	// 钩子
	Hook Hook
	// 自定义日记
	Logger Logger
	//
	userPool map[uint32]*User
}

// NewHub 创建一个枢纽
func NewHub(options ...Option) *Hub {
	opts := loadOptions(options...)
	h := &Hub{
		Hook:            opts.Hook,
		Logger:          opts.Logger,
		Persistence:     opts.Persistence,
		PersistenceType: opts.PersistenceType,
		userPool:        make(map[uint32]*User),
	}

	if h.Hook == nil {
		h.Hook = defaultHook
	}
	if h.Logger == nil {
		h.Logger = defaultLogger
	}

	return h
}

// AddUser 添加用户
func (h *Hub) AddUser(userId uint32) *User {
	// 判断用户是否存在
	if _, ok := h.userPool[userId]; !ok {
		// 添加用户
		h.userPool[userId] = &User{
			Id:         userId,
			ClientPool: make([]*Client, 0),
		}
	}
	return h.userPool[userId]
}

// AddClient 添加新的客户端
func (h *Hub) AddClient(userId uint32, ws *ws.Conn) *User {
	h.AddUser(userId).AddClient(ws)
	return h.userPool[userId]
}

// GetUser 获取用户
func (h *Hub) GetUser(userId uint32) *User {
	return h.userPool[userId]
}

// GetAllUser 获取所有用户列表
func (h *Hub) GetAllUser() map[uint32]*User {
	return h.userPool
}
