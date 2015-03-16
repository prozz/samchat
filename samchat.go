package main

import (
	"fmt"
	"github.com/prozz/mobster"
	"log"
	"os"
	"strconv"
)

func main() {
	s := mobster.NewServer()
	s.Debug = true

	s.OnConnect = func(ops *mobster.Ops, user, room string) {
		ops.SendToRoom(room, fmt.Sprintf("%s joins", user))
	}
	s.OnDisconnect = func(ops *mobster.Ops, user, room string) {
		ops.SendToRoom(room, fmt.Sprintf("%s leaves", user))
	}
	s.OnMessage = func(ops *mobster.Ops, user, room, message string) {
		ops.SendToRoom(room, fmt.Sprintf("%s: %s", user, message))
	}

	port, err := strconv.Atoi(os.Getenv("RUPPELLS_SOCKETS_LOCAL_PORT"))
	if err != nil {
		log.Fatalf("cannot get port: %s", err)
	}

	s.StartServerAndWait(port)
}
