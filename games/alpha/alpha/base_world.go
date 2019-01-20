package alpha

import (
	"fmt"
	"math"

	"github.com/SolarLune/resolv/resolv"
	"github.com/jacobrec/petah-royale/server/api"
	"github.com/jacobrec/petah-royale/server/core"
)

func StartWorld(gf core.GameIF) {

	g := makeGame(gf)
	fmt.Println("Started World")

	gf.OnJoin(func(id interface{}) { onJoin(&g, id) })
	gf.OnLeave(func(id interface{}) { onLeave(&g, id) })
	gf.OnEvent("shoot", func(id interface{}, ev api.Event) { onShoot(&g, id, ev) })
	gf.OnEvent("move", func(id interface{}, ev api.Event) { onMove(&g, id, ev) })
}

func makeGame(gf core.GameIF) gameObject {
	w := world{80, 60, make([]Moveable, 0), make([]Immoveable, 0)}
	//w.Walls = append(w.Walls, Immoveable{0,0,80,1}, Immoveable{0,0,1,60}, Immoveable{0,59,80,1}, Immoveable{79,0,1,60})
	walls, spawner := MakeMaze(6, 80, 60, .25)
	w.Walls = walls

	return gameObject{w, spawner, gf, make(map[interface{}]int, 3), make(map[int]interface{}, 3), 0}
}

// all movable objects are circles
type Moveable struct {
	Id     int     `json:"id"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Radius float64 `json:"radius"`
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

type Point struct {
	X float64
	Y float64
}

func onJoin(g *gameObject, id interface{}) {
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

func onShoot(g *gameObject, id interface{}, event api.Event) {
	shoot := event.Data.(*Shoot)
	p := getShotPath(g, shoot, g.connectionToGame[id])

	bang := Bang{shoot.X, shoot.Y, p.X, p.Y}
	ev := api.Event{"bang", bang}

	fmt.Println(p.X, p.Y)
	//distributeMessage(g, ev, id)
	g.gf.Send(ev, id)
}


func getShotPath(g *gameObject, shoot *Shoot, pid int) Point {
	var big = float64(g.w.Width * g.w.Height)
	shot := resolv.NewLine(int32(shoot.X*100), int32(shoot.Y*100), int32(shoot.X*100+big*math.Cos(shoot.Angle)), int32(shoot.Y*100+big*math.Sin(shoot.Angle)))

	var endX, endY int32
	for _, w := range g.w.Walls {
		wall := resolv.NewRectangle(int32(w.X*100), int32(w.Y*100), int32(w.Width*100), int32(w.Height*100))
		if shot.IsColliding(wall) {
            ps := shot.IntersectionPoints(wall)
            if len(ps) == 0{
                continue
            }
			p := shot.IntersectionPoints(wall)[0]
			if isP1Closer(int32(shoot.X*100), int32(shoot.Y*100), p.X, p.Y, endX, endY) {
				endX = p.X
				endY = p.Y
			}
		}
	}

    var hitplayer *Moveable
    for _, pl := range g.w.Players {
		pwall := resolv.NewRectangle(int32(pl.X*100), int32(pl.Y*100), int32(pl.Radius*2*100), int32(pl.Radius*2*100))
        if shot.IsColliding(pwall) && pl.Id != pid {
            fmt.Println("hello")
            ps := shot.IntersectionPoints(pwall)
            if len(ps) == 0 {
                continue
            }
			p := shot.IntersectionPoints(pwall)[0]
			if isP1Closer(int32(shoot.X*100), int32(shoot.Y*100), p.X, p.Y, endX, endY) {
				endX = p.X
				endY = p.Y
                hitplayer = &pl
			}
        }
    }

    if hitplayer != nil {
        fmt.Println("killed", hitplayer)
    }

	return Point{float64(endX) / 100, float64(endY) / 100}

}

func isP1Closer(x, y, x1, y1, x2, y2 int32) bool {
	return square(x-x1)+square(y-y1) < square(x-x2)+square(y-y2)
}

func square(x int32) int32 {
	return x * x
}
