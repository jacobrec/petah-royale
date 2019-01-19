package api

import (
    "fmt"
    "reflect"
    "encoding/json"
)

type rawEvent struct {
    Action string        `json:"action"`
    Data json.RawMessage `json:"data"`
}

type Event struct {
    Action string
    Data interface{}
}

type ActionReader struct {
    Intermap map[string]reflect.Type
}

func EmptyAR() ActionReader {
    return ActionReader{make(map[string]reflect.Type)}
}

func (ar *ActionReader) AddAction(key string, typ reflect.Type) {
    ar.Intermap[key] = typ
}

func (ar *ActionReader) Get(key string) reflect.Type {
    return ar.Intermap[key]
}

func (ar *ActionReader) ToAction(data []byte) Event {
    var evt rawEvent
    json.Unmarshal(data, &evt)

    fmt.Println(evt.Action)

    umdata := reflect.New(ar.Get(evt.Action)).Interface()
    json.Unmarshal(evt.Data, &umdata)

    return Event{evt.Action, umdata}
}
