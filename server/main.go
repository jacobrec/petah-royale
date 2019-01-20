package main

import (
    "github.com/jacobrec/petah-royale/server/server"
    "github.com/jacobrec/petah-royale/server/api"
)

func main() {
    wsgame := server.NewWSgame(api.DefaultEventsAR())
    server.Begin(wsgame, ":8049")
}
