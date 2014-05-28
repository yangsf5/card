// Author: sheppard(ysf1026@gmail.com) 2014-05-27

package user

import (
	"runtime/debug"
	"fmt"
	"github.com/yangsf5/claw/center"
)

func (u *User) EnterService(service string) {
	u.services.PushBack(service)
}

func (u *User) LeaveService(service string) {
	for e := u.services.Front(); e != nil; e = e.Next() {
		if e.Value == service {
			u.services.Remove(e)
			break
		}
	}
}

//TODO
func (u *User) LeaveAllService() {
	for e := u.services.Front(); e != nil; {
		center.Send("", e.Value.(string), u.sessionId, center.MsgTypeText, nil)
		u.services.Remove(e)
		fmt.Println(string(debug.Stack()))
	}
}
