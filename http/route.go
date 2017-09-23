package http

import (
	"net/http"
)

type Route struct {
	route   string
	handler http.Handler // 必须实现ServeHTTP方法，通过http.HandlerFunc可以实现转换
}
