package handler

import (
	"net/http"
	"time"

	"github.com/golang/net/websocket"
	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service/web"
)

func hallHandler(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		User string
	}
	userName := r.FormValue("user")
	glog.Infof("User enter hall, name=%s", userName)

	web.RenderHtml(w, "hall.html", &Param{userName})
}

func hallSocketHandler(ws *websocket.Conn) {
	userName := ws.Request().FormValue("user")
	glog.Infof("Hall socket, name=%s", userName)

	center.Send("Web", "CardWebsocket", 0, center.MsgTypeSystem, ws)

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

