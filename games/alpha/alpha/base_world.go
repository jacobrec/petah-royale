package alpha

import (
    "github.com/jacobrec/petah-royale/server/api"
    "github.com/jacobrec/petah-royale/server/core"
    "fmt"
)

func StartWorld(gf core.GameIF) {

    g := makeGame(gf)

    gf.OnJoin(func(id interface{}) { onJoin(&g, id) })
    gf.OnLeave(func(id interface{}) { onLeave(&g, id) })
    gf.OnEvent("shoot", func(id interface{}, ev api.Event) { onShoot(&g, id, ev) })
    gf.OnEvent("move", func(id interface{}, ev api.Event) { onMove(&g, id, ev) })
}

func makeGame(gf core.GameIF) gameObject {
    w := world{80, 60, make([]Moveable, 0), make([]Immoveable, 0)}
    //w.Walls = append(w.Walls, Immoveable{0,0,80,1}, Immoveable{0,0,1,60}, Immoveable{0,59,80,1}, Immoveable{79,0,1,60})
    walls, spawner := MakeMaze(5, 80, 60, .25)
    w.Walls = walls

    return gameObject{w, spawner, gf, make(map[interface{}]int, 3), make(map[int]interface{}, 3), 0}
}


// all movable objects are circles
type Moveable struct {
    Id int `json:"id"`
    X float64 `json:"x"`
    Y float64 `json:"y"`
    Radius float64 `json:"radius"`
}

// all immovable objects are rectangles
type Immoveable struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
    Width float64 `json:"width"`
    Height float64 `json:"height"`
}

type world struct {
    Width int `json:"width"`
    Height int `json:"height"`
    Players []Moveable `json:"players"`
    Walls []Immoveable `json:"walls"`
}

func (w* world) getPlayerById(id int) *Moveable {
    for _, p := range w.Players {
        if p.Id == id {
            return &p
        }
    }
    return nil
}

type gameObject struct {
    w world
    spawner func() (float64, float64)
    gf core.GameIF
    connectionToGame map[interface{}]int
    gameToConnection map[int]interface{}
    LastId int
}

type InitialMessage struct {
    Id int `json:"id"`
    X float64 `json:"x"`
    Y float64 `json:"y"`
    World world `json:"world"`
}


func onJoin(g *gameObject, id interface{}){
    pid := g.LastId
    g.LastId++

    g.connectionToGame[id] = pid
    g.gameToConnection[pid] = id

    x, y := g.spawner()
    player := Moveable{pid, x, y, 0.5}
    g.w.Players = append(g.w.Players, player)

    data := New{player.Id, player.X, player.Y, player.Radius}
    ev := api.Event{"new", data}
    distributeMessage(g, ev, id)

    initData := InitialMessage{pid, player.X, player.Y, g.w}
    g.gf.Send(api.Event{"initial", initData}, id)

}

func onLeave(g *gameObject, id interface{}){
    pid := g.connectionToGame[id]
    data := Exit{pid}
    ev := api.Event{"exit", data}
    distributeMessage(g, ev, id)
}

func onMove(g *gameObject, id interface{}, event api.Event){
    move := event.Data.(*Move)
    pid := g.connectionToGame[id]

    pp := g.w.getPlayerById(pid)
    pp.X = move.X
    pp.Y = move.Y

    data := Draw{pp.Id, pp.X, pp.Y}
    ev := api.Event{"draw", data}
    distributeMessage(g, ev, id)
}


func distributeMessage(g* gameObject, ev api.Event, not interface{}){
    for _, conn := range g.gameToConnection {
        if not != conn {
            g.gf.Send(ev, conn)
        }
    }
}

func onShoot(gf *gameObject, id interface{}, event api.Event){
    shoot := event.Data.(*Shoot)
    fmt.Println("BANG")
    fmt.Println(shoot.Angle)
}
