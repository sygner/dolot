package main

import (
	"dolott_game/internal/server"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	server.RunServer()
}
