package core

import (
	"github.com/jacobrec/petah-royale/server/api"
)

type GameIF interface {
	OnEvent(action string, cback func(id interface{}, evt api.Event))
	OnJoin(cback func(id interface{}))
	OnLeave(cback func(id interface{}))
	Send(evt api.Event, id interface{})
	GetClients() map[interface{}]bool
	Disconnect(id interface{})
}

// GameMux provides utilities for GameIFs
type GameMux struct {
	ar       api.ActionReader
	clients  map[interface{}]bool
	handlers map[string]func(id interface{}, evt api.Event)
	joinfun  func(id interface{})
	leavefun func(id interface{})
}

func (g *GameMux) OnEvent(action string, cback func(id interface{}, evt api.Event)) {
	g.handlers[action] = cback
}

func (g *GameMux) OnJoin(cback func(id interface{})) {
	g.joinfun = cback
}

func (g *GameMux) OnLeave(cback func(id interface{})) {
	g.leavefun = cback
}

func (g *GameMux) HandleEvt(data []byte, id interface{}) {
	evt := g.ar.ToAction(data)
	g.handlers[evt.Action](id, evt)
}

func (g *GameMux) HandleJoin(id interface{}) {
	g.clients[id] = true
	g.joinfun(id)
}

func (g *GameMux) HandleLeave(id interface{}) {
	delete(g.clients, id)
	g.leavefun(id)
}
