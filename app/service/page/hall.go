// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package page

import (
	"net/http"

	"github.com/golang/glog"
)

func hallHandler(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		User string
	}
	userName := r.FormValue("user")
	glog.Infof("User enter hall, name=%s", userName)

	renderHtml(w, "hall.html", &Param{userName})
}
