package main

import (
    "github.com/jacobrec/petah-royale/server/server"
    "github.com/jacobrec/petah-royale/games/alpha/alpha"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "math/rand"
    "time"
)


func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    router := gin.Default()

    config := cors.DefaultConfig()
	config.AllowAllOrigins = true
    router.Use(cors.New(config))

    server.CreateGameServer(&router.RouterGroup, func() *server.WSgame {
        wsgame := server.NewWSgame(alpha.EventsAR())
        alpha.StartWorld(&wsgame)
        return &wsgame
    })

    router.Run(":8049")
}
