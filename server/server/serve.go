package server

import (
    "github.com/gin-gonic/gin"
)


func Begin(wsgame *WSgame, port string) {
    router := gin.Default()

    router.GET("/wsgame", func(c *gin.Context) {
        wsgame.GameHandler(c.Writer, c.Request)
    })

    router.Run(port)
}
