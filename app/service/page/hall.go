// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package page

import (
	"net/http"

	"code.google.com/p/go.net/websocket"
	"github.com/golang/glog"

	"github.com/yangsf5/card/app/logic/door"
)

func hallHandler(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		User string
	}
	userName := r.FormValue("user")
	glog.Infof("User enter hall, name=%s", userName)

	renderHtml(w, "hall.html", &Param{userName})
}

func hallSocketHandler(ws *websocket.Conn) {
	userName := ws.Request().FormValue("user")
	glog.Infof("Hall socket, name=%s", userName)

	door.Login(ws, userName)
}

