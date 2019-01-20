package core

import (
    "github.com/jacobrec/petah-royale/server/api"
)

func StartWorld(gf *GameIF) {

    g := makeGame(gf)

    gf.OnJoin(func(id interface{}) { onJoin(&g, id) })
    gf.OnLeave(func(id interface{}) { onLeave(&g, id) })
    gf.OnEvent("shoot", func(id interface{}) { onShoot(&g, id) })
    gf.OnEvent("move", func(id interface{}) { onMove(&g, id) })
}

func makeGame(gf *GameIF) gameObject {
    w := world{80, 60, make([]moveable, 0), make([]immoveable, 0)}
    w.walls = append(w.walls, immoveable{0,0,80,1}, immoveable{0,0,1,60}, immoveable{0,59,80,1}, immoveable{79,0,1,60})

    return gameObject{w, gf}
}


// all movable objects are circles
type moveable struct {
    id int
    x float64
    y float64
    radius float64
}

// all immovable objects are rectangles
type immoveable struct {
    x float64
    y float64
    width float64
    height float64
}

type world struct {
    width int
    height int
    players moveable[]
    walls immovable[]
}

func (w* world) getPlayerById(id) *moveable {
    for _, p := range w.players {
        if p.id == p {
            return &p
        }
    }
}

type gameObject struct {
    w world
    gf *GameIF
    connectionToGame[interface{}] int
    gameToConnection[int] interface{}
}

var playerCount int

func onJoin(g *gameObject, id interface{}){
    pid = playerCount
    playerCount++

    connectionToGame[id] = pid
    gameToConnection[pid] = id

    player := moveable{pid, 2, 2, 0.5}
    g.w.players = append(g.w.players, player)

    g.sendPlayerMove(player)
}

func onLeave(g *gameObject, id interface{}){
}

func onMove(g *gameObject, id interface{}, event api.Event){
    move := event.Data.(api.Move)
    pid := connectionToGame[id]

    pp := g.w.getPlayerById(pid)
    pp.x = move.X
    pp.y = move.Y

    g.sendPlayerMove(*pp)
}

func (g* gameObject) sendPlayerMove(player moveable){
    data := api.Draw{player.id, player.x, player.y}
    ev := api.Event{"draw", data}
    for _, conn := range g.gameToConnection {
        g.gf.Send(ev, conn)
    }
}

func onShoot(gf *gameObject, id interface{}, event api.Shoot){
    shoot := event.Data.(api.Shoot)
}
