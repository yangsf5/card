// Author: sheppard(ysf1026@gmail.com) 2014-05-23

package service

import (
	"code.google.com/p/go.net/websocket"
	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"

	"github.com/yangsf5/card/app/logic"
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
			go logic.Login(conn)
		}
	}
}

func (s* Websocket) ClawStart() {
	s.conns = make(map[int]*websocket.Conn)
}

func (s* Websocket) AddClient() {
}
