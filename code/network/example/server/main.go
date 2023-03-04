package main

import "game/myrepo/code/network"

func main() {
	server := network.NewServer(":8023", "tcp")
	server.Run()
	select {}

}
