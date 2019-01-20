package server

import (
    "github.com/jacobrec/petah-royale/server/api"
    "github.com/jacobrec/petah-royale/server/core"
    "net/http"
    "github.com/gorilla/websocket"
    "encoding/json"
    "fmt"
)


var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type WSgame struct {
    core.GameMux
}

func NewWSgame(ar api.ActionReader) WSgame {
    return WSgame{core.NewGameMux(ar)}
}

func (wsg *WSgame) Send(evt api.Event, id interface{}) {
    conn := id.(*websocket.Conn)
    msg, _ := json.Marshal(evt)
    conn.WriteMessage(websocket.TextMessage, msg)
}

func (wsg *WSgame) Disconnect(id interface{}) {
    conn := id.(*websocket.Conn)
    conn.Close()
}


func (wsg *WSgame) GameHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    wsg.HandleJoin(conn)

    for {
        t, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }

        wsg.HandleEvt(msg, &conn)
        conn.WriteMessage(t, msg)
    }

    wsg.HandleLeave(&conn)
}
