package routes

import (
	lilyhttp "github.com/hxlhxl/go-utils/http"
)

type User struct {
	lilyhttp.RouteHandler
}

func (h *User) GET() {

	type UD struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	ud := &UD{Name: "huaxiong", Age: 90}
	h.JsonData["json"] = ud
	// h.ServeJSON()

}

func init() {
	lilyhttp.InitRoute("/api/v1/user", &User{})
}
