// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package main

import (
	"time"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"

	"github.com/yangsf5/card/app/handler"
	myService "github.com/yangsf5/card/app/service"
)

var (
)

func main() {
	glog.Info("Card start!")

	service.Register()
	myService.Register()

	handler.RegisterHandler()

	center.Use([]string{"Web", "CardWebsocket"})

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Card exit!")
	glog.Flush()
}

