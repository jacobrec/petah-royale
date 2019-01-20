package main

import (
    "github.com/jacobrec/petah-royale/server/server"
    "github.com/jacobrec/petah-royale/server/api"
    "github.com/jacobrec/petah-royale/server/core"
)

func main() {
    wsgame := server.NewWSgame(api.DefaultEventsAR())
    core.StartWorld(&wsgame)
    server.Begin(&wsgame, ":8049")
}
