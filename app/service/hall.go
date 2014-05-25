// Author: sheppard(ysf1026@gmail.com) 2014-05-24

package service

import (
	"fmt"

	"github.com/yangsf5/card/app/service/hall"
)

type Hall struct {
}

func (s* Hall) ClawCallback(session int, source string, msgType int, msg interface{}) {
}

func (s* Hall) ClawStart() {
}


//TODO
type HallUser struct {
}
