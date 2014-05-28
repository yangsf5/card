// Author: sheppard(ysf1026@gmail.com) 2014-05-24

package service

import (
	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"

	"github.com/yangsf5/card/app/service/hall"
)

type Hall struct {
}

func (s* Hall) ClawCallback(session int, source string, msgType int, msg interface{}) {
	glog.Infof("Service.CardHall recv type=%v msg=%v", msgType, msg)
	switch msgType {
	case center.MsgTypeSystem:
		if user, ok := msg.(hall.User); ok {
			if ret := hall.Enter(session, user); !ret {
				glog.Info("Service.CardHall enter hall failed")
				return
			}
			user.Login()
			go user.Tick()
		} else {
			glog.Info("Service.CardHall msg is not a net.Peer")
		}
	case center.MsgTypeText:
		if msg, ok := msg.(string); ok {
			if msg == "LEAVE" {
				hall.Leave(session)
			}
		}
	}
}

func (s* Hall) ClawStart() {
}

