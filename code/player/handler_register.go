package player

import (
	"fmt"
	"game/myrepo/code/chat"
	"game/myrepo/code/commonFunction"
)

type Handler func(interface{})

func (p *Player) AddFriend(data interface{}) {
	fId := data.(uint64)
	if !commonFunction.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

func (p *Player) DelFriend(data interface{}) {
	fID := data.(uint64)
	p.FriendList = commonFunction.DelEleInSlice(fID, p.FriendList)
}

func (p *Player) ResolveChatMsg(data interface{}) {
	msg := data.(chat.Msg)
	fmt.Println(msg)
	// todo
}
