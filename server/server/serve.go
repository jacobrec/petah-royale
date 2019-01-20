package server

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func getTest(c *gin.Context) {
    c.String(http.StatusOK, "test string")
}

func Begin(wsgame WSgame, port string) {
    router := gin.Default()

    router.GET("/test", getTest)
    router.GET("/wsgame", func(c *gin.Context) {
        wsgame.GameHandler(c.Writer, c.Request)
    })
    router.Run(port)
}
