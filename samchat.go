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

	s.OnConnect = func(ctx *mobster.Ops, name, room string) {
		ctx.SendToRoom(room, fmt.Sprintf("%s joins", name))
	}
	s.OnDisconnect = func(ctx *mobster.Ops, name, room string) {
		ctx.SendToRoom(room, fmt.Sprintf("%s leaves", name))
	}
	s.OnMessage = func(ctx *mobster.Ops, name, room, message string) {
		ctx.SendToRoom(room, fmt.Sprintf("%s: %s", name, message))
	}

	s.StartServerAndWait(*port)
}
