package alpha

import (
    "reflect"
    "github.com/jacobrec/petah-royale/server/api"
)

func EventsAR() api.ActionReader {
    return api.ActionReader{map[string]reflect.Type{
        "move": reflect.TypeOf(Move{}),
        "shoot": reflect.TypeOf(Shoot{}),

        "draw": reflect.TypeOf(Draw{}),
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

