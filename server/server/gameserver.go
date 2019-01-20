package server

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "math/rand"
)


type GameServer struct {
    gamefun func() *WSgame
    games map[string]*WSgame
}

func NewGameServer(gamefun func() *WSgame) GameServer {
    return GameServer{gamefun, make(map[string]*WSgame)}
}

func CreateGameServer(router *gin.RouterGroup, gamefun func() *WSgame) {
    gs := NewGameServer(gamefun)
    gs.Route(router)
}

func (gm *GameServer) PostCreate(c *gin.Context) {
    id := randString(32)
    gm.games[id] = gm.gamefun()
    c.String(http.StatusOK, id)
}

func (gm *GameServer) GameWS(c *gin.Context) {
    id := c.Params.ByName("id")
    game := gm.games[id]

    if game == nil {
        c.String(http.StatusNotFound, id)
        return
    }

    game.GameHandler(c.Writer, c.Request)
}

func (gm *GameServer) DelGame(c *gin.Context) {
    id := c.Params.ByName("id")
    delete(gm.games, id)
    c.String(http.StatusOK, id)
}

func (gm *GameServer) Route(router *gin.RouterGroup) {
    router.POST("/new", gm.PostCreate)
    router.GET("/game/:id", gm.GameWS)
    router.DELETE("/game/:id", gm.DelGame)
}


const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

