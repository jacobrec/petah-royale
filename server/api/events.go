package api

import (
    "reflect"
)

func DefaultEventsAR() ActionReader {
    return ActionReader{map[string]reflect.Type{
        "move": reflect.TypeOf(Move{}),
        "shoot": reflect.TypeOf(Shoot{}),
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
