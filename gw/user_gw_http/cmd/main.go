package main

import "dolott_user_gw_http/internal/server"

func main() {
	// fmt.Println(float32(math.Round(float64(0.2) * constants.TICKET_BUY_RATE)))
	server.RunServer()
}
