// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package service

import (
	"github.com/yangsf5/claw/center"
)


func Register() {
	services := map[string]center.Service{
		"CardWebsocket": &Websocket{},
	}

	for name, cb := range services {
		center.Register(name, cb)
	}
}

func send(source, destination string, session int, msgType int, msg interface{}) {
	center.Send(source, destination, session, msgType, msg)
}
