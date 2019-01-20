package server

import (
    "github.com/jacobrec/petah-royale/server/api"
    "github.com/jacobrec/petah-royale/server/core"
    "net/http"
    "github.com/gorilla/websocket"
    "encoding/json"
    "fmt"
    "sync"
)


var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

type WSgame struct {
    core.GameMux
}

type concConn struct {
    conn *websocket.Conn
    mu sync.Mutex
}

func NewWSgame(ar api.ActionReader) WSgame {
    return WSgame{core.NewGameMux(ar)}
}

func (wsg *WSgame) Send(evt api.Event, id interface{}) {
    conn := id.(*concConn)
    msg, _ := json.Marshal(evt)
    conn.mu.Lock()
    defer conn.mu.Unlock()
    conn.conn.WriteMessage(websocket.TextMessage, msg)
}

func (wsg *WSgame) Disconnect(id interface{}) {
    conn := id.(*concConn)
    conn.mu.Lock()
    defer conn.mu.Unlock()
    conn.conn.Close()
}


func (wsg *WSgame) GameHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := wsupgrader.Upgrade(w, r, nil)
    cconn := new(concConn)
    cconn.conn = conn

    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    wsg.HandleJoin(cconn)

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }

        wsg.HandleEvt(msg, cconn)
    }

    wsg.HandleLeave(cconn)
}
