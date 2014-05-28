// Author: sheppard(ysf1026@gmail.com) 2014-05-28

package hall

import (
	"github.com/golang/glog"

	"github.com/yangsf5/card/app/logic/proto"
)

func HandleClientMessage(session int, msgType string, msgData interface{}) {
	u, ok := sessions[session]
	if !ok {
		glog.Errorf("HallClientMessage session not in hall, sessionId=%d", session)
		return
	}

	switch msgType {
	case "chat":
		chatMsg := &proto.HCChat{u.Name(), msgData.(string)}
		Broadcast(proto.Encode("Cardhall", chatMsg))
	}
}
