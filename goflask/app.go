package goflask

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Ip          string
	Port        string
	namedRoutes map[string]*Route
}

func NewApp() *App {
	return &App{Ip: "127.0.0.1", Port: "9999", namedRoutes: make(map[string]*Route)}
}
func (app *App) addHandler(url string, handler http.Handler) {
	app.namedRoutes[url] = &Route{handler: handler}
}
func (app *App) Route(url string, handler func(http.ResponseWriter, *http.Request)) *App {
	// 使其具备ServeHTTP能力
	app.addHandler(url, http.HandlerFunc(handler))
	return app
}

// func (app *App)  Method(methods string) *App {

// }

func (app *App) mux(w http.ResponseWriter, r *http.Request) (http.Handler, error) {
	reqURL := r.URL.Path
	if route, ok := app.namedRoutes[reqURL]; ok {
		return route.handler, nil
	} else {
		return nil, errors.New("无法MUXING")
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 如果是模板路由怎么办
	// 如果是静态服务器怎么办
	// 如果是API怎么办
	h, err := app.mux(w, r)
	if err == nil {
		h.ServeHTTP(w, r)
	}
}

func (app *App) Start() {
	addr := app.Ip + ":" + app.Port
	fmt.Println("http app starting", addr)
	// 不使用DefaultServeMux派发路由
	log.Fatal("ListenAndServe Fail:", http.ListenAndServe(addr, app))
}
func (app *App) serveStatic() {

}

func (app *App) ServeStatic() {
	// 对URL进行判断

	// viewsStaticHandler := http.FileServer(http.Dir(filepath.Join(Root, "/views")))
	// staticStaticHanlder := http.FileServer(http.Dir(filepath.Join(Root, "/debug")))
	// // 存在问题，下面多个都会有派发，导致multiple writerHeader
	// stripPrefix("/debug", staticStaticHanlder).ServeHTTP(w, r)

	// viewsStaticHandler.ServeHTTP(w, r)
}
