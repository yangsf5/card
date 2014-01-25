// Author: sheppard(ysf1026@gmail.com) 2014-01-19

package net

type User interface {
	Send(string)
}

type Group struct {
	users map[string] User

	broadcast chan string
}

func NewGroup() *Group {
	g := &Group{}
	g.users = make(map[string] User)
	g.broadcast = make(chan string)

	go g.tick()
	return g
}

func (g *Group) AddUser(uid string, u User) bool {
	if _, ok := g.users[uid]; ok {
		return false
	}
	g.users[uid] = u
	return true
}

func (g *Group) DelUser(uid string) {
	delete(g.users, uid)
}

func (g *Group) GetUsers() map[string] User{
	return g.users
}

func (g *Group) Broadcast(msg string) {
	g.broadcast <- msg
}

func (g *Group) tick() {
	for {
		msg := <-g.broadcast
		for _, u := range g.users {
			u.Send(msg)
		}
	}
}

