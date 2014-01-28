// Author: sheppard(ysf1026@gmail.com) 2014-01-25

package room

import (
//	"github.com/yangsf5/card/app/engine/net"
)

var (
	fightRoom *FightRoom
)

func init() {
	fightRoom = NewFightRoom()
}

/*
func Enter(roomType string, uid string, u net.User) Room {
	switch roomType {
	case "pvp":
		if user, ok := u.(FightUser); ok {
			fightRoom.Enter(uid, user)
			return &fightRoom
		}
	}
	return nil
}
*/

func EnterFightRoom(uid string, u FightUser) *FightRoom {
	fightRoom.Enter(uid, u)
	return fightRoom
}
