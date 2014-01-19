// Author: sheppard(ysf1026@gmail.com) 2014-01-08

package room

import (
	"github.com/yangsf5/card/app/logic/user"
)

func init() {
}

type Room struct {
	Users map[string] *user.User
}

func NewRoom() *Room {
	room := &Room{}
	room.Users = make(map[string] *user.User)
	return room
}

func (r *Room) AddUser(u *user.User) {
	if u == nil {
		return
	}
	if _, ok := r.Users[u.Name]; ok {
		return
	}
	r.Users[u.Name] = u
}

func (r *Room) DelUser(u *user.User) {
	if u == nil {
		return
	}
	delete(r.Users, u.Name)
}

func (r *Room) Tick() {
	for {
	}
}

