// Author: sheppard(ysf1026@gmail.com) 2014-01-08

package room

import (
	"github.com/yangsf5/card/app/engine/net"
	"github.com/yangsf5/card/app/logic/user"
)

func init() {
}

type FightRoom struct {
	net.Group
}

func NewFightRoom() *FightRoom {
	room := &FightRoom{}
	return room
}

func (r *FightRoom) Enter(uid string, u net.User) bool {
	ret := r.AddUser(uid, u)
	if ret {
	}
	return ret
}

func (r *FightRoom) Leave(uid string) bool {
	r.DelUser(uid)
}
