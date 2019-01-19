package server

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func getTest(c *gin.Context) {
    c.String(http.StatusOK, "test string")
}

func Begin(port string) {

    router := gin.Default()

    router.GET("/test", getTest)
    router.GET("/ws", func(c *gin.Context) {
        wshandler(c.Writer, c.Request)
    })
    router.Run(port)
}
