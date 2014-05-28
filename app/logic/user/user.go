// Author: sheppard(ysf1026@gmail.com) 2013-12-25

package user

import (
	"container/list"
	"fmt"
	"github.com/yangsf5/card/app/logic/proto"
	"github.com/yangsf5/card/app/logic/room"
	"github.com/yangsf5/card/app/service/hall"
)

type User struct {
	sessionId int
	name string
	RecvMsg <-chan string
	SendMsg chan<- string
	Offline <-chan error

	disconnected bool

	curRoom *room.FightRoom
	enemy room.FightUser

	services *list.List
}

func NewUser(sessionId int, name string, recv <-chan string, send chan<- string, offline <-chan error) *User {
	u := &User{}
	u.sessionId = sessionId
	u.name, u.RecvMsg, u.SendMsg, u.Offline = name, recv, send, offline
	u.services = list.New()
	return u
}

func (u *User) SessionId() int {
	return u.sessionId
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Tick() {
	for {
		select {
		case msg, ok := <-u.RecvMsg:
			if !ok {
				u.Logout("Recv channel closed")
				return
			}
			pack := proto.Decode(msg)
			fmt.Println("User recv:", pack.Type, pack.Data)
			u.handle(pack.Service, pack.Type, pack.Data)
		case err, ok := <-u.Offline:
			if !ok {
				u.Logout("Offline channel closed")
			} else {
				u.Logout(err.Error())
			}
			return
		}
	}
}

func (u *User) Send(msg []byte) {
	if !u.disconnected {
		u.SendMsg <- string(msg)
		fmt.Println("User send:", string(msg))
	}
}

func (u *User) Login() {
	fmt.Println("User login, name:", u.name)
}

func (u *User) Logout(reason string) {
	if !u.disconnected {
		u.disconnected = true
		close(u.SendMsg)
		u.LeaveAllService()
		if u.curRoom != nil {
			u.curRoom.Leave(u.name)
		}
		fmt.Println("User disconneted, err:", reason)
	}
}

func (u *User) handle(service, msgType string, msgData interface{}) {
	switch msgType {
	case "chat":
		chatMsg := &proto.HCChat{u.name, msgData.(string)}
		hall.Broadcast(proto.Encode("Cardhall", chatMsg))
	case "enterRoom":
		switch msgData.(string) {
		case "pvp":
			u.curRoom = room.EnterFightRoom(u.name, u)
		case "pve":
		}
	}
}

func (u *User) SetEnemy(enemy room.FightUser) {
	u.enemy = enemy
}

func (u *User) GetEnemy() room.FightUser {
	return u.enemy
}

