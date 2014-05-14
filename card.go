// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package main

import (
	"time"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"

	myService "github.com/yangsf5/card/app/service"
)

var (
)

func main() {
	glog.Info("Card start!")

	service.Register()
	myService.Register()
	center.Use([]string{"CardHttp"})

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Card exit!")
	glog.Flush()
}

