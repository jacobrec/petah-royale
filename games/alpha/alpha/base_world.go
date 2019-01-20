package alpha

import (
	"fmt"
	"math"

	"github.com/jacobrec/petah-royale/server/api"
	"github.com/jacobrec/petah-royale/server/core"
)

func StartWorld(gf core.GameIF) {

    TestRectorsector()

	g := makeGame(gf)
	fmt.Println("Started World")

	gf.OnJoin(func(id interface{}) { onJoin(&g, id) })
	gf.OnLeave(func(id interface{}) { onLeave(&g, id) })
	gf.OnEvent("shoot", func(id interface{}, ev api.Event) { onShoot(&g, id, ev) })
	gf.OnEvent("move", func(id interface{}, ev api.Event) { onMove(&g, id, ev) })
}

func makeGame(gf core.GameIF) gameObject {
	w := world{80, 60, make([]Moveable, 0), make([]Immoveable, 0)}
	walls, spawner := MakeMaze(6, w.Width, w.Height, .25)
	w.Walls = walls

	return gameObject{w, spawner, gf, make(map[interface{}]int, 3), make(map[int]interface{}, 3), 0}
}

// all movable objects are circles
type Moveable struct {
	Id     int     `json:"id"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Size float64 `json:"radius"`
}

// all immovable objects are rectangles
type Immoveable struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type world struct {
	Width   int          `json:"width"`
	Height  int          `json:"height"`
	Players []Moveable   `json:"players"`
	Walls   []Immoveable `json:"walls"`
}

func (w *world) getPlayerById(id int) *Moveable {
	for _, p := range w.Players {
		if p.Id == id {
			return &p
		}
	}
	return nil
}

type gameObject struct {
	w                world
	spawner          func() (float64, float64)
	gf               core.GameIF
	connectionToGame map[interface{}]int
	gameToConnection map[int]interface{}
	LastId           int
}

func onJoin(g *gameObject, id interface{}) {
	pid := g.LastId
	g.LastId++

	g.connectionToGame[id] = pid
	g.gameToConnection[pid] = id

	x, y := g.spawner()
	player := Moveable{pid, x, y, 0.5}
	g.w.Players = append(g.w.Players, player)

	data := New{player.Id, player.X, player.Y, player.Size/2}
	ev := api.Event{"new", data}
	distributeMessage(g, ev, id)

	initData := InitialMessage{pid, player.X, player.Y, g.w}
	g.gf.Send(api.Event{"initial", initData}, id)

}

func onLeave(g *gameObject, id interface{}) {
	pid := g.connectionToGame[id]
	data := Exit{pid}
	ev := api.Event{"exit", data}
	distributeMessage(g, ev, id)

	removePlayer(g, id)
}

func removePlayer(g *gameObject, id interface{}) {
	var na []Moveable
	var pid = g.connectionToGame[id]
	for _, v := range g.w.Players {
		if v.Id == pid {
			continue
		} else {
			na = append(na, v)
		}
	}
	g.w.Players = na
}

func onMove(g *gameObject, id interface{}, event api.Event) {
	move := event.Data.(*Move)
	pid := g.connectionToGame[id]

	pp := g.w.getPlayerById(pid)
	pp.X = move.X
	pp.Y = move.Y

	data := Draw{pp.Id, pp.X, pp.Y}
	ev := api.Event{"draw", data}
	distributeMessage(g, ev, id)
}

func distributeMessage(g *gameObject, ev api.Event, not interface{}) {
	for _, conn := range g.gameToConnection {
		if not != conn {
			g.gf.Send(ev, conn)
		}
	}
}
func sendToAll(g *gameObject, ev api.Event) {
	for _, conn := range g.gameToConnection {
        g.gf.Send(ev, conn)
	}
}

func onShoot(g *gameObject, id interface{}, event api.Event) {
	shoot := event.Data.(*Shoot)
	p := getShotPath(g, shoot, g.connectionToGame[id])

	bang := Bang{shoot.X, shoot.Y, p.x, p.y}
	ev := api.Event{"bang", bang}

	fmt.Println(p.x, p.y)
	sendToAll(g, ev)
}


func getShotPath(g *gameObject, shoot *Shoot, pid int) Point {
	var big = float64(g.w.Width * g.w.Height)
	shot := LineSeg{Point{shoot.X, shoot.Y}, Point{shoot.X+big*math.Cos(shoot.Angle), shoot.Y+big*math.Sin(shoot.Angle)}}

	var endX, endY float64
	for _, w := range g.w.Walls {
		wall := Rectangle{w.X, w.Y, w.Width, w.Height}
		if IsRectorsect(wall, shot) {
            p, _ := Rectorsect(wall, shot)
			if isP1Closer(shoot.X, shoot.Y, p.x, p.y, endX, endY) {
				endX = p.x
				endY = p.y
			}
		}
	}
	shot = LineSeg{Point{shoot.X, shoot.Y}, Point{endX, endY}}

    fmt.Println("WALL", endX, endY)
    fmt.Println("Checking players")
    var hitplayer Moveable
    hitplayer.Id = -1
    for _, pl := range g.w.Players {
		pwall := Rectangle{pl.X, pl.Y, pl.Size, pl.Size}
        if IsRectorsect(pwall, shot) && pl.Id != pid {
            p, _ := Rectorsect(pwall, shot)
			if isP1Closer(shoot.X, shoot.Y, p.x, p.y, endX, endY) {
				endX = p.x
				endY = p.y
                hitplayer = pl
                fmt.Println(hitplayer)
			}
        }
    }

    if hitplayer.Id != -1 {
        fmt.Println("killed", hitplayer, "by", pid)
        dead := Dead{hitplayer.Id}
        ev := api.Event{"dead", dead}
        sendToAll(g, ev)
        g.gf.Disconnect(g.gameToConnection[hitplayer.Id])
    }

	return Point{endX, endY}
}

func isP1Closer(x, y, x1, y1, x2, y2 float64) bool {
	return square(x-x1)+square(y-y1) < square(x-x2)+square(y-y2)
}

func square(x float64) float64 {
	return x * x
}
