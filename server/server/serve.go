package server

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jacobrec/petah-royale/server/api"
)

func getTest(c *gin.Context) {
    c.String(http.StatusOK, "test string")
}

func Begin(port string) {

    wsgame := NewWSgame(api.DefaultEventsAR())

    router := gin.Default()

    router.GET("/test", getTest)
    router.GET("/wsgame", func(c *gin.Context) {
        wsgame.GameHandler(c.Writer, c.Request)
    })
    router.Run(port)
}
