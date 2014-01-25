// Author: sheppard(ysf1026@gmail.com) 2014-01-25

package room

import (
	"github.com/yangsf5/card/app/engine/net"
)

var (
	fightRoom FightRoom
)

func init() {
	//fightRoom FightRoom
}

func Enter(roomType string, uid string, u net.User) *FightRoom {
	switch roomType {
	case "pvp":
		fightRoom.Enter(uid, u)
		return &fightRoom
	}
	return nil
}
