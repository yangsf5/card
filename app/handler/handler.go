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
