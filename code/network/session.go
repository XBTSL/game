package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Session struct {
	conn    net.Conn
	packer  *NormalPacker
	chWrite chan *Message
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPack(binary.BigEndian),
		chWrite: make(chan *Message, 1),
	}
}

func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

func (s *Session) Read() {
	err := s.conn.SetReadDeadline(time.Now().Add(time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	for {
		message, err := s.packer.Unpack(s.conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("server receive message:", string(message.Data))
		s.chWrite <- &Message{
			Id:   999,
			Data: []byte("hi tomjack"),
		}
	}
}

func (s *Session) Write() {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case msg := <-s.chWrite:
			s.send(msg)
		}
	}
}

func (s *Session) send(message *Message) {
	byte, err := s.packer.Pack(message)
	if err != nil {
		return
	}
	_, err = s.conn.Write(byte)
	if err != nil {
		fmt.Println(err)
	}
}
