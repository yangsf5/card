// Author: sheppard(ysf1026@gmail.com) 2014-01-12

package hall

import (
	"fmt"
	"github.com/yangsf5/card/app/engine/net"
	"github.com/yangsf5/card/app/logic/proto"
)

var (
	group *net.Group
)

func init() {
	group = net.NewGroup()
}

func Enter(uid string, u net.User) bool {
	ret := group.AddUser(uid, u)
	if ret {
		msg := &proto.HCRoomList{}
		for k, _ := range configs {
			protoRoom := proto.Room{k, 0, fmt.Sprintf("/room/default?room=%s&user=%s", k, uid)}
			msg.Rooms = append(msg.Rooms, protoRoom)
		}

		fmt.Println(msg)
		u.Send(proto.Encode(msg))
	}
	return ret
}

func Broadcast(msg string) {
	group.Broadcast(msg)
}

func DelUser(uid string) {
	group.DelUser(uid)
}
