// Author: sheppard(ysf1026@gmail.com) 2013-12-25

package user

import (
	"fmt"
	"github.com/yangsf5/card/app/logic/hall"
	"github.com/yangsf5/card/app/logic/proto"
	"github.com/yangsf5/card/app/logic/room"
)

type User struct {
	Name string
	RecvMsg <-chan string
	SendMsg chan<- string
	Offline <-chan error

	disconnected bool

	curRoom *room.FightRoom
	enemy room.FightUser
}

func NewUser(name string, recv <-chan string, send chan<- string, offline <-chan error) *User {
	u := &User{}
	u.Name, u.RecvMsg, u.SendMsg, u.Offline = name, recv, send, offline
	return u
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
			u.handle(pack.Type, pack.Data)
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

func (u *User) Send(msg string) {
	if !u.disconnected {
		u.SendMsg <- msg
		fmt.Println("User send:", msg)
	}
}

func (u *User) Login() {
	fmt.Println("User login, name:", u.Name)
}

func (u *User) Logout(reason string) {
	if !u.disconnected {
		u.disconnected = true
		close(u.SendMsg)
		hall.DelUser(u.Name)
		if u.curRoom != nil {
			u.curRoom.DelUser(u.Name)
		}
		fmt.Println("User disconneted, err:", reason)
	}
}

func (u *User) handle(msgType string, msgData interface{}) {
	switch msgType {
	case "chat":
		chatMsg := &proto.HCChat{u.Name, msgData.(string)}
		hall.Broadcast(proto.Encode(chatMsg))
	case "enterRoom":
		switch msgData.(string) {
		case "pvp":
			u.curRoom = room.EnterFightRoom(u.Name, u)
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
