// Author: sheppard(ysf1026@gmail.com) 2014-05-14

package page

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"runtime/debug"

	"code.google.com/p/go.net/websocket"
	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
)

const (
	TEMPLATE_DIR = "./view"
)

var (
	templates = make(map[string] *template.Template)
)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	checkError(err)

	header := TEMPLATE_DIR + "/common/header.html"
	footer := TEMPLATE_DIR + "/common/footer.html"
	glog.Infof("Common template [%s %s]", header, footer)

	var tplName, tplPath string
	for _, fileInfo := range fileInfoArr {
		tplName = fileInfo.Name()
		if ext := path.Ext(tplName); ext != ".html" {
			continue
		}
		tplPath = TEMPLATE_DIR + "/" + tplName
		glog.Infof("Loading template %s", tplPath)
		t := template.Must(template.ParseFiles(tplPath, header, footer))
		templates[tplPath] = t
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = TEMPLATE_DIR + "/" + tmpl
	tpl, ok := templates[tmpl];
	if !ok {
		glog.Errorf("Render html, but template is nil, name=%s", tmpl)
		return
	}

	err := tpl.Execute(w, data)
	checkError(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "index.html", nil)
	}
	if r.Method == "POST" {
		userName := r.PostFormValue("user")
		if userName != "" {
			glog.Infof("User login, name=%s", userName)
			http.Redirect(w, r, "/hall/hall?user="+userName, http.StatusFound)
		} else {
			renderHtml(w, "index.html", nil)
			glog.Infof("Index post user is empty")
		}
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		http.ServeFile(w, r, file)
	})
}

func Start() {
	type HttpConfig struct {
		ListenAddr string `xml:"listenAddr,attr"`
	}
	type ConfigPack struct {
		XMLName xml.Name `xml:"clawconfig"`
		Http HttpConfig `xml:"http"`
	}
	var cfg ConfigPack
	center.GetConfig(&cfg)

	mux := http.NewServeMux()
	staticDirHandler(mux, "/public/js/", "./public/js/")
	mux.HandleFunc("/", safeHandler(indexHandler))
	mux.HandleFunc("/hall/hall", safeHandler(hallHandler))
	mux.Handle("/hall/hall/socket", websocket.Handler(hallSocketHandler))
	err := http.ListenAndServe(cfg.Http.ListenAddr, mux)
	checkError(err)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Errorf("WARN: panic in %v - %v", fn, err)
				glog.Errorf(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
