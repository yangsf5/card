// Author: sheppard(ysf1026@gmail.com) 2014-01-25

package room

import (
	"github.com/yangsf5/card/app/engine/net"
)

type FightUser interface {
	net.User
	GetLevel() int
}

type FightRoom struct {
	Room
}

func (r *FightRoom) Start(uid string) net.User {
	/*
	for k, v := range r.GetUsers() {
		if _, ok := v.(FightUser); ok && k != uid {
			//TODO
			return v
		}
	}
	*/
	return nil
}
