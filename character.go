package main

type clist struct {
	Characters []character `json:"characters,[]characters"`
}
type character struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
    Class string `json:"class"`
    Race string `json:"race"`
    Coin int `json:"coin"`
    Exp int `json:"exp"`
    Attributes map[string]int `json:"attributes"`
    Savings map[string]int `json:"savings"`
}
