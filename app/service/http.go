// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package service

import (
	"github.com/golang/glog"

	"github.com/yangsf5/card/app/service/page"
)

type Http struct {
}

func (s *Http) ClawCallback(session int, source string, msgType int, msg interface{}) {
}

func (s *Http) ClawStart() {
	glog.Infof("Http service start")
	page.Start()
}

