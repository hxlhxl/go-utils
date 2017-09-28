package http

import (
	"fmt"
)

type ApiWrapper interface {
	OPTIONS()
	HEAD()
	GET()

	ServeJSON() interface{}
	// POST()
	// PUT()
	// DELETE()
}

type RouteHandler struct {
	JsonData map[interface{}]interface{}
}

func (c *RouteHandler) Init() {
	c.JsonData = make(map[interface{}]interface{})
}

func (c *RouteHandler) OPTIONS() {
	fmt.Println("RouteHandler OPTIONS")
}
func (c *RouteHandler) HEAD() {
	fmt.Println("RouteHandler HEAD")

}
func (c *RouteHandler) GET() {
	fmt.Println("RouteHandler GET")
}

func (c *RouteHandler) ServeJSON() interface{} {
	return c.JsonData["json"]
}

// func (h *RouteHandler) GET() {
// 	fmt.Fprintf(h.w, fmt.Sprintf("%v", h.JsonData))
// }
