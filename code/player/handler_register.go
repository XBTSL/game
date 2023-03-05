package player

import (
	"fmt"
	"game/myrepo/code/commonFunction"
	"game/myrepo/code/network"
	"game/myrepo/proto/protogo/player"
	"google.golang.org/protobuf/proto"
)

type Handler func(packet *network.SessionPacket)

func (p *Player) AddFriend(packet *network.SessionPacket) {
	req := &player.C2SAddFriend{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	if !commonFunction.CheckInNumberSlice(req.Uid, p.FriendList) {
		p.FriendList = append(p.FriendList, req.Uid)
	}
}

func (p *Player) DelFriend(packet *network.SessionPacket) {
	req := &player.C2SDelFriend{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	p.FriendList = commonFunction.DelEleInSlice(req.Uid, p.FriendList)
}

func (p *Player) ResolveChatMsg(packet *network.SessionPacket) {
	req := &player.C2SSendChatMsg{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)
	// todo
}
