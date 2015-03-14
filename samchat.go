package main

import (
	"github.com/prozz/mobster"
	"fmt"
	"flag"
)

func main() {
	port := flag.Int("port", 8765, "chat server port")
	flag.Parse()

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

	s.StartServerAndWait(*port)
}
