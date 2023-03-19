package gchat

type Hub struct {
	Option   *Options
	UserPool map[int]*User
}

// NewHub 创建一个枢纽
func NewHub(options ...Option) *Hub {
	opts := loadOptions(options...)
	return &Hub{
		Option:   opts,
		UserPool: make(map[int]*User),
	}
}

// AddUser 添加用户
func (h *Hub) AddUser(userId int) *User {
	// 判断用户是否存在
	if _, ok := h.UserPool[userId]; !ok {
		// 添加用户
		h.UserPool[userId] = &User{
			Id:         userId,
			ClientPool: make(map[int]*Client),
		}
	}
	return h.UserPool[userId]
}

// AddClient 添加连接客户端
func (h *Hub) AddClient(userId int, client *Client) *User {
	// 判断用户是否存在
	if _, ok := h.UserPool[userId]; !ok {
		// 添加用户
		h.UserPool[userId] = &User{
			Id:         0,
			ClientPool: make(map[int]*Client),
		}
	}
	// 用户添加连接
	h.UserPool[userId].ClientPool[client.Id] = client
	return h.UserPool[userId]
}
