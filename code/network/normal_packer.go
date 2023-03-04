package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

type NormalPacker struct {
	Order binary.ByteOrder
}

func NewNormalPack(order binary.ByteOrder) *NormalPacker {
	return &NormalPacker{
		order,
	}
}

// pack | data 长度| id| data
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	buffer := make([]byte, 8+8+len(message.Data))
	p.Order.PutUint64(buffer[:8], uint64(len(message.Data)))
	p.Order.PutUint64(buffer[8:16], message.Id)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

// unpack
func (p *NormalPacker) Unpack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}
	length := p.Order.Uint64(buffer[:8])
	id := p.Order.Uint64(buffer[8:])
	dataLenth := length - 16
	dataBuffer := make([]byte, dataLenth)
	_, err = io.ReadFull(reader, dataBuffer)
	if err != nil {
		return nil, err
	}
	msg := &Message{
		Id:   id,
		Data: dataBuffer,
	}
	return msg, nil
}
