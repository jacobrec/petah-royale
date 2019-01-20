package main

import (
    "github.com/jacobrec/petah-royale/server/server"
    "github.com/jacobrec/petah-royale/server/api"
    "github.com/jacobrec/petah-royale/server/core"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    server.CreateGameServer(&router.RouterGroup, func() *server.WSgame {
        wsgame := server.NewWSgame(api.DefaultEventsAR())
        core.StartWorld(&wsgame)
        return &wsgame
    })

    router.Run(":8049")
}
