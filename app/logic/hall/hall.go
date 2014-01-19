// Author: sheppard(ysf1026@gmail.com) 2014-01-12

package hall

import (
	"fmt"
	"github.com/yangsf5/card/app/logic/proto"
)

type User interface {
	Send(string)
}

var (
	users map[string] User

	broadcast chan string
)

func init() {
	users = make(map[string] User)
	broadcast = make(chan string)

	go Tick()
}

func Enter(uid string, u User) bool {
	ret := AddUser(uid, u)
	if ret {
		msg := &proto.HCRoomList{}
		for k, _ := range configs {
			protoRoom := proto.Room{k, 0, fmt.Sprintf("/room/default?room=%s&user=%s", k, uid)}
			msg.Rooms = append(msg.Rooms, protoRoom)
		}

		fmt.Println(msg)
		u.Send(proto.Encode(msg))
	}
	return ret
}

func AddUser(uid string, u User) bool {
	if _, ok := users[uid]; ok {
		return false
	}
	users[uid] = u
	return true
}

func DelUser(uid string) {
	delete(users, uid)
}

func Broadcast(msg string) {
	broadcast <- msg
}

func Tick() {
	for {
		msg := <-broadcast
		fmt.Println("Hall broadcast msg:", msg)
		for _, u := range users {
			u.Send(msg)
		}
	}
}

