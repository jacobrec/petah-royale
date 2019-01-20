package server


type GameServer struct {
    gamefun func() *WSgame
    games map[string]*WSgame
}

func NewGameServer(gamefun func() core.GameIF) {
    return GameServer{gamefun, make(map[string]core.GameIF}
}

func (gm *GameServer) PostCreate(c *gin.Context) {
    id := randString(32)
    gm.games[id] = gamefun()
    c.String(http.StatusOK, id)
}

func (gm *GameServer) GameWS(c *gin.Context) {
    id := c.Params.ByName("id")
    game := games[id]

    if game == nil {
        c.String(http.StatusNotFound, id)
        return
    }

    game.GameHandler(c.Writer, c.Request)
}


const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

