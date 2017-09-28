package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/hxlhxl/go-utils/path"
)

type ShellServer struct {
	Ip         string
	Port       string
	StaticPath map[string]string
	Router     map[string]*Route
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

type UD struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// func InitRoute(uri string, handler http.HandlerFunc) {
// func InitRoute(uri string, rh *RouteHandler) {
func InitRoute(uri string, aw ApiWrapper) {
	// Routes = append(Routes, &Route{uri, http.HandlerFunc(handler)})
	// http.HandlerFunc就是把一个handler转换为拥有ServeHttp方法的接口类型
	// reflect接口是value,type 对

	reflectVal := reflect.ValueOf(aw)
	t := reflect.Indirect(reflectVal).Type()
	fmt.Println("reflectVal", reflectVal, "reflect type", t)

	ctl := reflect.New(t)
	init := ctl.MethodByName("Init")
	init.Call(nil)

	h := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			method := ctl.MethodByName("GET")
			fmt.Println("controller type", ctl.Type(), "method name", runtime.FuncForPC(method.Pointer()).Name(), "method type", method.Type(), "method string", method.String())
			// 调用controller中的GET|POST等方法
			method.Call(nil)

			// fmt.Println(ctl.Kind(), ctl.Elem())
			// elem := ctl.Elem()
			// fmt.Println("jsondata", elem.FieldByName("JsonData"))
			// jsonData := elem.FieldByName("JsonData")

			// js, _ := jsonData.Interface()

			// if js, ok := jsonData["json"]; ok {
			// 	fmt.Println("解析json成功！", js)
			// }
			// jsonData := elem.FieldByName("JsonData").Interface()
			// fmt.Println("interface jsondata", jsonData, jsonData.json)

			get_json := ctl.MethodByName("ServeJSON")
			js := get_json.Call(nil)
			fmt.Println("%v", js)
			x, err := json.Marshal(js)
			// ud := &UD{Name: "huaxiong", Age: 24}
			// ujs, err := json.Marshal(ud)
			if err != nil {
				fmt.Println("%v", err)
			}
			if err == nil {
				fmt.Println("marshal json %v", x)
				w.Header().Set("Content-Type", "application/json")
				w.Write(x)
				return
				// fmt.Fprintf(w, fmt.Sprintf("%s", []byte(jsonData)))
			}
		}

	}
	Routes[uri] = &Route{uri, http.HandlerFunc(h)}
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
