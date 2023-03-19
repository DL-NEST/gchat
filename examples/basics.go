package main

import (
	"fmt"
	"gchat"
	"github.com/gin-gonic/gin"
	"net/http"
	ws "nhooyr.io/websocket"
)

func main() {
	g := gchat.NewHub(gchat.WithPersistence(true))
	gin.SetMode(gin.ReleaseMode)
	srv := gin.Default()

	srv.GET("/socket", func(ctx *gin.Context) {
		Conn, err := ws.Accept(ctx.Writer, ctx.Request, &ws.AcceptOptions{InsecureSkipVerify: true})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		Conn.CloseRead(ctx)
		//g.AddUser(896).AddClient(0,Conn)
		g.AddClient(242, &gchat.Client{
			Id: 0,
			Ws: Conn,
		})
		//defer SConn.Close(ws.StatusNormalClosure, "")
	})
	srv.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, g.UserPool)
	})

	err2 := srv.Run(":9011")
	if err2 != nil {
		fmt.Println("服务启动失败")
		return
	}
}
