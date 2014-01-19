// Author: sheppard(ysf1026@gmail.com) 2014-01-18

package controllers

import (
	"github.com/robfig/revel"
)

type Table struct {
	*revel.Controller
}

func (c Table) Default(room string, table string, user string) revel.Result {
	return c.Render(room, table, user)
}
