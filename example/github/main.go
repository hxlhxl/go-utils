package main

import (
	//"fmt"
	
	_ "go-labs/github/routes"

	"github.com/hxlhxl/go-utils/http"
)

func main() {
	ip := "127.0.0.1"
	port := "9999"
	httpShell := http.InitHttpShell(ip, port)
	httpShell.Start()
}
