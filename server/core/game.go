package core

import (
    "github.com/jacobrec/petah-royale/server/api"
)

type GameIF interface {
    OnEvent(action string, cback func(g *GameIF, id interface{}, evt api.Event))
    OnJoin(cback func(g *GameIF, id interface{}))
    OnLeave(cback func(g *GameIF, id interface{}))
    Send(evt api.Event, id interface{})
    GetClients() map[interface{}]bool
    Disconnect(id interface{})
}


// GameMux provides utilities for GameIFs
type GameMux struct {
    ar api.ActionReader
    clients map[interface{}]bool
    handlers map[string]func(g *GameIF, id interface{}, evt api.Event)
    joinfun func(g *GameIF, id interface{})
    leavefun func(g *GameIF, id interface{})
}

func (g *GameMux) OnEvent(action string, cback func(gif *GameIF, id interface{}, evt api.Event)) {
    g.handlers[action] = cback
}

func (g *GameMux) OnJoin(cback func(gif *GameIF, id interface{})) {
    g.joinfun = cback
}

func (g *GameMux) OnLeave(cback func(gif *GameIF, id interface{})) {
    g.leavefun = cback
}

func (g *GameMux) HandleEvt(gif *GameIF, data []byte, id interface{}) {
    evt := g.ar.ToAction(data)
    g.handlers[evt.Action](gif, id, evt)
}

func (g *GameMux) HandleJoin(gif *GameIF, id interface{}) {
    g.clients[id] = true
    g.joinfun(gif, id)
}

func (g *GameMux) HandleLeave(gif *GameIF, id interface{}) {
    delete(g.clients, id)
    g.leavefun(gif, id)
}
