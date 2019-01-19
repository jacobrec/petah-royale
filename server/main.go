package main

import (
    "fmt"
    "reflect"
    _ "github.com/jacobrec/petah-royale/server/server"
    "github.com/jacobrec/petah-royale/server/api"
)

type A struct {
    A int `json:"a"`
    B int `json:"b"`
}

type B struct {
    B int `json:"b"`
}

func main() {
    //server.Begin(":8049")
    ar := api.EmptyAR()

    ar.AddAction("A", reflect.TypeOf(A{}))
    ar.AddAction("B", reflect.TypeOf(B{}))

    fmt.Println(ar.ToAction([]byte(`{"action": "A", "data": {"a": 5, "b": 6}}`)).Data.(*A))
    fmt.Println(ar.ToAction([]byte(`{"action": "B", "data": {"b": 6}}`)).Data.(*B))
}
