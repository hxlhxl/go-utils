package http

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/hxlhxl/go-utils/path"
)

type ShellServer struct {
	Ip         string
	Port       string
	StaticPath map[string]string
	Router     map[string]*Route
}

type HandlerWrapper struct {
	method string
}

// var Routes []*Route
var Routes map[string]*Route = make(map[string]*Route)
var Root string

func InitHttpShell(ip, port string) *ShellServer {
	Root = path.Getwd()
	routerPath := Root + "/routes"
	isExist := path.IsExist(routerPath)
	if !isExist {
		log.Fatalln("读取HTTP服务的路由路径失败")
	}
	HttpShell := &ShellServer{
		Ip:         ip,
		Port:       port,
		StaticPath: make(map[string]string),
		Router:     Routes,
	}
	return HttpShell
}

func InitRoute(uri string, handler http.HandlerFunc) {
	// Routes = append(Routes, &Route{uri, http.HandlerFunc(handler)})
	// http.HandlerFunc就是把一个handler转换为拥有ServeHttp方法的接口类型
	Routes[uri] = &Route{uri, http.HandlerFunc(handler)}
}

func stripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			h.ServeHTTP(w, r2)
		}
	})
}

func (shell *ShellServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// /api/v1/user/...
	value, ok := shell.Router[r.URL.Path]
	if ok {
		value.handler.ServeHTTP(w, r)
		// shell.Router.router.handler.ServeHTTP(w,r)
		return
	}

	// 静态路由
	// /index.html	/files/a.txt
	viewsStaticHandler := http.FileServer(http.Dir(filepath.Join(Root, "/views")))
	staticStaticHanlder := http.FileServer(http.Dir(filepath.Join(Root, "/debug")))
	// 存在问题，下面多个都会有派发，导致multiple writerHeader
	stripPrefix("/debug", staticStaticHanlder).ServeHTTP(w, r)

	viewsStaticHandler.ServeHTTP(w, r)
	// 交给DefaultServeMux如何派发路由给不同的handler
	// http.Handle("/", http.StripPrefix("/views/", viewsStaticHandler))
	// http.HandleFunc根据路由叫给不同的handler，只不过这个handler是通过以上的HanlderFunc转换具有ServeHTTP

	// http.NotFound(w, r)
	return
}

func (shell *ShellServer) InitStaticPath(uri string, path string) {
	shell.StaticPath[uri] = path
}

func (shell *ShellServer) Start() error {
	addr := shell.Ip + ":" + shell.Port
	fmt.Println("http shell starting", addr)
	// 不使用DefaultServeMux派发路由
	log.Fatal("ListenAndServe Fail:", http.ListenAndServe(addr, shell))
	return nil
}
