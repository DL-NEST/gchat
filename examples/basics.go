package main

import (
	"context"
	"fmt"
	"gchat"
	"github.com/gin-gonic/gin"
	"net/http"
	ws "nhooyr.io/websocket"
	"time"
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
		//g.AddUser(896).AddClient(0,Conn)
		//ct, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Printf("关闭\n")
			Conn.Close(ws.StatusNormalClosure, "")
		}()
		go func(c context.Context) {
			fmt.Println("监听")
			_, _, errRead := Conn.Reader(c)
			errStart := ws.CloseStatus(errRead)
			fmt.Printf("errStart:%s\n", errStart)
			switch errStart {
			case ws.StatusNormalClosure:
				fmt.Println("正常退出")
			}
			fmt.Println("监听结束")
			return
		}(ctx)
		fmt.Println("结束")
	})
	srv.GET("/list", func(ctx *gin.Context) {

		ctx.AbortWithStatus(http.StatusMultiStatus)
		ctx.JSON(http.StatusOK, g.GetAllUser())
	})

	err2 := srv.Run(":9011")
	if err2 != nil {
		fmt.Println("服务启动失败")
		return
	}
}
