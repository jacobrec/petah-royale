package api

import (
    "reflect"
)

func DefaultEventsAR() ActionReader {
    return ActionReader{map[string]reflect.Type{
        "move": reflect.TypeOf(Move{}),
    }}
}


type Move struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}
