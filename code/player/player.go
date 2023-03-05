package player

import (
	"game/myrepo/code/network"
	"game/myrepo/proto/protogo/messageId"
)

type Player struct {
	Uid            uint64
	FriendList     []uint64
	HandlerParamCh chan *network.SessionPacket
	handlers       map[messageId.MessageId]Handler
	session        *network.Session
}

func NewPlayer() *Player {
	p := &Player{
		Uid:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[messageId.MessageId]Handler),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[messageId.MessageId(handlerParam.Msg.Id)]; ok {
				fn(handlerParam)
			}
		}
	}
}
