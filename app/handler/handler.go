// Author: sheppard(ysf1026@gmail.com) 2014-05-19

package handler

import (
	"github.com/yangsf5/claw/service/web"
)

func RegisterHandler() {
	web.StaticDirHandler("/public/js/", "./public/js/")
	web.RegisterHttpHandler("/", indexHandler)
	web.RegisterHttpHandler("/hall/hall", hallHandler)
	web.RegisterWebSocketHandler("/hall/hall/socket", hallSocketHandler)
}
