package main

import "game/myrepo/code/network"

func main() {
	client := network.NewClient(":8023")
	client.Run()
	select {}
}
