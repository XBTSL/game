package manager

import "game/myrepo/code/player"

// 维护在线玩家
type PlayerMgr struct {
	player map[uint64]player.Player
	addPch chan player.Player
}

func (pm *PlayerMgr) Add(p player.Player) {
	pm.player[p.Uid] = p
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPch:
			pm.Add(p)
		}
	}
}
