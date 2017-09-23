package http

import (
	"net/http"
)

type ApiWrapper interface {
	OPTIONS()
	HEAD()
	GET()
	POST()
	PUT()
	DELETE()
}

type RouteHandler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *RouteHandler) GET() {

}
