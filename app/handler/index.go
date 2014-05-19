// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package handler

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/service/web"
)

const (
	TEMPLATE_DIR = "./view"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		web.RenderHtml(w, "index.html", nil)
	}
	if r.Method == "POST" {
		userName := r.PostFormValue("user")
		if userName != "" {
			glog.Infof("User login, name=%s", userName)
			http.Redirect(w, r, "/hall/hall?user="+userName, http.StatusFound)
		} else {
			web.RenderHtml(w, "index.html", nil)
			glog.Infof("Index post user is empty")
		}
	}
}

