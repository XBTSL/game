package player

import "game/myrepo/proto/protogo/messageId"

func (p *Player) HandlerRegister() {
	p.handlers[messageId.MessageId_C2SAddFriend] = p.AddFriend
	p.handlers[messageId.MessageId_C2SDelFriend] = p.DelFriend
	p.handlers[messageId.MessageId_S2CSendFriend] = p.ResolveChatMsg
}
