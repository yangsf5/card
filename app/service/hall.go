package service

import (
	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"

	"github.com/yangsf5/card/app/service/hall"
	"github.com/yangsf5/card/app/logic/proto"
)

type Hall struct {
}

func (s* Hall) ClawCallback(session int, source string, msgType int, msg interface{}) {
	glog.Infof("Service.CardHall recv session=%d type=%v msg=%v", session, msgType, msg)
	switch msgType {
	case center.MsgTypeSystem:
		if user, ok := msg.(hall.User); ok {
			if ret := hall.Enter(session, user); !ret {
				glog.Info("Service.CardHall enter hall failed")
				user.Kick("Repeated login")
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
	case center.MsgTypeClient:
		if pack, ok := msg.(*proto.Pack); ok {
			hall.HandleClientMessage(session, pack.Type, pack.Data)
		} else {
			glog.Errorf("Service.CardHall MsgTypeClient msg is not a *proto.Pack")
		}
	}
}

func (s* Hall) ClawStart() {
}

