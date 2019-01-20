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
    ar api.ActionReader
    clients map[interface{}]bool
    handlers map[string]func(id interface{}, evt api.Event)
    joinfun func(id interface{})
    leavefun func(id interface{})
}

func NewGameMux(ar api.ActionReader) GameMux {
    return GameMux{
        ar,
        make(map[interface{}]bool),
        make(map[string]func(id interface{}, evt api.Event)),
        nil,
        nil,
    }
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

func (g *GameMux) GetClients() map[interface{}]bool {
    return g.clients
}

func (g *GameMux) HandleEvt(data []byte, id interface{}) {
    evt := g.ar.ToAction(data)
    handler := g.handlers[evt.Action]
    if handler != nil {
        handler(id, evt)
    }
}

func (g *GameMux) HandleJoin(id interface{}) {
    g.clients[id] = true
    if g.joinfun != nil {
        g.joinfun(id)
    }
}

func (g *GameMux) HandleLeave(id interface{}) {
    delete(g.clients, id)
    if g.leavefun != nil {
        g.leavefun(id)
    }
}
