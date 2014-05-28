// Author: sheppard(ysf1026@gmail.com) 2014-01-12

package hall

import (
	"fmt"
	"github.com/yangsf5/claw/engine/net"
	"github.com/yangsf5/card/app/logic/proto"
)

type User interface {
	net.Peer
	Name() string

	//TODO remove
	Login()
	Tick()

	EnterService(service string)
}

var (
	group *net.Group
	sessions map[int]User
)

func init() {
	group = net.NewGroup()
	sessions = make(map[int]User)
}

func Enter(session int, u User) bool {
	ret := group.AddPeer(u.Name(), u)
	if ret {
		u.EnterService("CardHall")
		sessions[session] = u
		msg := &proto.HCRoomList{}
		for k, _ := range configs {
			protoRoom := proto.Room{k, 0, fmt.Sprintf("#")}
			msg.Rooms = append(msg.Rooms, protoRoom)
		}

		fmt.Println(msg)
		u.Send([]byte(proto.Encode(msg)))
	}
	return ret
}

func Leave(session int) {
	if u, ok := sessions[session]; ok {
		group.DelPeer(u.Name())
		delete(sessions, session)
	}
}

func Broadcast(msg string) {
	group.Broadcast([]byte(msg))
}

