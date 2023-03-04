package main

import "game/myrepo/code/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}
