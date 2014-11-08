package main

import (
    "math/rand"
)

type id [2]uint64

type Entity struct {
    ID id
}

func NewID() id {
    return id{uint64(rand.Int63()), uint64(rand.Int63())}
}

