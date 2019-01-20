package alpha

import (
    "reflect"
    "github.com/jacobrec/petah-royale/server/api"
)

func EventsAR() api.ActionReader {
    return api.ActionReader{map[string]reflect.Type{
        "move": reflect.TypeOf(Move{}),
        "shoot": reflect.TypeOf(Shoot{}),

        "new": reflect.TypeOf(New{}),
        "exit": reflect.TypeOf(Exit{}),
        "draw": reflect.TypeOf(Draw{}),
        "bang": reflect.TypeOf(Bang{}),
        "dead": reflect.TypeOf(Dead{}),
    }}
}


type Shoot struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
    Angle float64 `json:"angle"`
    Weapon int `json:"weapon"`
}

type Move struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type Draw struct {
    Id int `json:"id"`
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type New struct {
    Id int `json:"id"`
    X float64 `json:"x"`
    Y float64 `json:"y"`
    Size float64 `json:"radius"`
}

type Exit struct {
    Id int `json:"id"`
}

type Bang struct {
    X1 float64 `json:"x1"`
    Y1 float64 `json:"y1"`
    X2 float64 `json:"x2"`
    Y2 float64 `json:"y2"`
}

type Dead struct {
    Id int `json:"id"`
}

type InitialMessage struct {
    Id int `json:"id"`
    X float64 `json:"x"`
    Y float64 `json:"y"`
    World world `json:"world"`
}
