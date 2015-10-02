package user

import (
	"container/list"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"

	"github.com/yangsf5/card/app/logic/proto"
	"github.com/yangsf5/card/app/logic/room"
)

type User struct {
	sessionId int
	name string
	RecvMsg <-chan string
	SendMsg chan<- string
	recvErr, sendErr <-chan error

	disconnected bool

	curRoom *room.FightRoom
	enemy room.FightUser

	services *list.List
}

func NewUser(sessionId int, name string, recv <-chan string,  send chan<- string, recvErr, sendErr <-chan error) *User {
	u := &User{}
	u.sessionId = sessionId
	u.name, u.RecvMsg, u.SendMsg = name, recv, send
	u.recvErr, u.sendErr = recvErr, sendErr
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
				u.RecvMsg = nil
				break
			}
			pack := proto.Decode(msg)
			glog.Infof("User recv, service=%s type=%s data=%v", pack.Service, pack.Type, pack.Data)
			//u.handle(pack.Service, pack.Type, pack.Data)
			center.Send("", pack.Service, u.sessionId, center.MsgTypeClient, pack)
		case err, _ := <-u.recvErr:
			u.Logout(err.Error())
			return
		case err, _ := <-u.sendErr:
			u.Logout(err.Error())
			return
		}
	}
}

func (u *User) Send(msg []byte) {
	if !u.disconnected {
		u.SendMsg <- string(msg)
		glog.Infof("User send, msg=%v", string(msg))
	}
}

func (u *User) Login() {
	glog.Infof("User login, name=%s", u.name)
}

func (u *User) Logout(reason string) {
	if !u.disconnected {
		u.disconnected = true
		close(u.SendMsg)
		u.LeaveAllService()
		if u.curRoom != nil {
			u.curRoom.Leave(u.name)
		}
		glog.Infof("User disconneted, reason=%s", reason)
	}
}

func (u *User) Kick(reason string) {
	u.Logout(reason)
}

//TODO remove
func (u *User) handle(service, msgType string, msgData interface{}) {
	switch msgType {
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

