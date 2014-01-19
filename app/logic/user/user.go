// Author: sheppard(ysf1026@gmail.com) 2013-12-25

package user

import (
	"fmt"
	"github.com/yangsf5/card/app/logic/hall"
	"github.com/yangsf5/card/app/logic/proto"
)

type User struct {
	Name string
	RecvMsg <-chan string
	SendMsg chan<- string
	Offline <-chan error

	disconnected bool
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
			fmt.Println("User recv:", msg)
			chatMsg := &proto.HCChat{u.Name, msg}
			hall.Broadcast(proto.Encode(chatMsg))
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
		fmt.Println("User disconneted, err:", reason)
	}
}
