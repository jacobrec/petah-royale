package core

import (
    "github.com/jacobrec/petah-royale/server/api"
)

func StartWorld(gf *GameIF) {

    g := makeGame(gf)

    gf.OnJoin(func(id interface{}) { onJoin(g, id) })
    gf.OnLeave(func(id interface{}) { onLeave(g, id) })
    gf.OnEvent("shoot", func(id interface{}) { onShoot(g, id) })
    gf.OnEvent("move", func(id interface{}) { onMove(g, id) })
}

func makeGame(gf *GameIF) {
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

type gameObject struct {
    w world
    gf *GameIF
}

func onJoin(gf gameObject, id interface{}){
}

func onLeave(gf gameObject, id interface{}){
}

func onMove(gf gameObject, id interface{}, event api.Event){
    move := event.Data.(api.Move)
}

func onShoot(gf gameObject, id interface{}, event api.Shoot){
    shoot := event.Data.(api.Shoot)
}
