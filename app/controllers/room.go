// Author: sheppard(ysf1026@gmail.com) 2014-01-14

package controllers

import (
	"github.com/robfig/revel"
)

type Room struct {
	*revel.Controller
}

func (c Room) Default(room string, user string) revel.Result {
	return c.Render(room, user)
}
