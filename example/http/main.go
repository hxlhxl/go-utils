package main

import (
	_ "github.com/hxlhxl/go-utils/example/http/routes"
	"github.com/hxlhxl/go-utils/http"
)

func main() {
	ip := "127.0.0.1"
	port := "9999"
	httpShell := http.InitHttpShell(ip, port)
	httpShell.Start()
}
