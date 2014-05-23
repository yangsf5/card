// Author: sheppard(ysf1026@gmail.com) 2014-05-23

package service

import (
	"code.google.com/p/go.net/websocket"
	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"

	"github.com/yangsf5/card/app/logic/hall"
	"github.com/yangsf5/card/app/logic/user"
)

type Websocket struct {
	sessionIdGenerator int
	conns map[int]*websocket.Conn
}

func (s* Websocket) ClawCallback(session int, source string, msgType int, msg interface{}) {
	glog.Infof("Service.Websocket recv type=%v msg=%v", msgType, msg)
	switch msgType {
	case center.MsgTypeSystem:
		if conn, ok := msg.(*websocket.Conn); ok {
			s.sessionIdGenerator += 1
			s.conns[s.sessionIdGenerator] = conn

			s.login(conn)
		}
	}
}

func (s* Websocket) ClawStart() {
	s.conns = make(map[int]*websocket.Conn)
	hall.ReadConfig("config/room.xml")
}

func (s* Websocket) AddClient() {
}

func (s* Websocket) login(conn *websocket.Conn) {
	userName := conn.Request().FormValue("user")

	offline := make(chan error)

	recvMsg := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(conn, &msg)
			if err != nil {
				close(recvMsg)
				return
			}
			recvMsg <- msg
		}
	}()

	sendMsg := make(chan string)
	go func() {
		for {
			select {
			case msg, ok := <-sendMsg:
				// If the channel is closed, they disconnected.
				if !ok {
					return
				}

				if err := websocket.Message.Send(conn, msg); err != nil {
					// Disconneted.
					offline <- err
					return
				}
			}
		}
	}()

	u := user.NewUser(userName, recvMsg, sendMsg, offline)
	ok := hall.Enter(u.Name, u)
	if !ok {
		return
	}
	u.Login()
	go u.Tick()
}

