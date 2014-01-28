// Author: sheppard(ysf1026@gmail.com) 2014-01-25

package room

import (
	"github.com/yangsf5/card/app/engine/net"
	"github.com/yangsf5/card/app/logic/proto"
)

type FightUser interface {
	net.User
	//GetLevel() int
	SetEnemy(FightUser)
	GetEnemy() FightUser
}

type FightRoom struct {
	Name string
	*net.Group
}

func NewFightRoom() *FightRoom {
	room := &FightRoom {
		"pvp",
		net.NewGroup(),
	}
	return room
}

func (r *FightRoom) Enter(uid string, u FightUser) bool {
	if(r.AddUser(uid, u)) {
		u.Send(proto.Encode(&proto.HCEnter{r.Name}))
		return true
	}
	return false
}

func (r *FightRoom) Leave(uid string) {
	r.DelUser(uid)
}

func (r *FightRoom) Start(uid string) {
	_, enemy := r.Find(func(id string, user net.User) bool {
		fightUser, ok := user.(FightUser);
		return ok && id != uid && fightUser.GetEnemy() == nil
	})

	if user := r.GetUser(uid); user != nil {
		user.(FightUser).SetEnemy(enemy.(FightUser))
	}

	return
}
