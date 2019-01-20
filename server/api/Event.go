package api

type Event struct {
    Action string        `json:"action"`
    Data interface{}     `json:"data"`
}
