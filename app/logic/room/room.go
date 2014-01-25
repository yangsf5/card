// Author: sheppard(ysf1026@gmail.com) 2014-01-08

package room

import (
	"github.com/yangsf5/card/app/engine/net"
)

func init() {
}

type Room struct {
	net.Group
}

func NewRoom() *Room {
	room := &Room{}
	return room
}

func (r *Room) Enter(uid string, u net.User) bool {
	ret := r.AddUser(uid, u)
	if ret {
	}
	return ret
}

func (r *Room) Leave(uid string) {
	r.DelUser(uid)
}
