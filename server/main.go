package main

import (
    "github.com/jacobrec/petah-royale/server/server"
    "github.com/jacobrec/petah-royale/server/api"
    "github.com/jacobrec/petah-royale/server/core"
    "github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
    router := gin.Default()
    router.Use(CORSMiddleware())

    server.CreateGameServer(&router.RouterGroup, func() *server.WSgame {
        wsgame := server.NewWSgame(api.DefaultEventsAR())
        core.StartWorld(&wsgame)
        return &wsgame
    })

    router.Run(":8049")
}
