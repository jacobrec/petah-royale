package main

import (
    "github.com/jacobrec/petah-royale/server/server"
    "github.com/jacobrec/petah-royale/server/api"
    "github.com/jacobrec/petah-royale/games/alpha/alpha"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)


func main() {
    router := gin.Default()

    config := cors.DefaultConfig()
	config.AllowAllOrigins = true
    router.Use(cors.New(config))

    server.CreateGameServer(&router.RouterGroup, func() *server.WSgame {
        wsgame := server.NewWSgame(api.DefaultEventsAR())
        alpha.StartWorld(&wsgame)
        return &wsgame
    })

    router.Run(":8049")
}
