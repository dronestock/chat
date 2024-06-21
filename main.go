package main

import (
	"github.com/dronestock/chat/internal"
	"github.com/dronestock/drone"
)

func main() {
	drone.New(internal.New).Boot()
}
