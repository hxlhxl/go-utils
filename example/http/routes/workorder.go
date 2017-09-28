package routes

import (
	"fmt"

	lilyhttp "github.com/hxlhxl/go-utils/http"
)

// 写法很丑，还需要内置的http包，非常的不好看

type Workorder struct {
	lilyhttp.RouteHandler
}

func (h *Workorder) GET() {
	fmt.Println("workorder")

}

// post方法,get_user,delete_user
func init() {
	lilyhttp.InitRoute("/api/v1/workorder", &Workorder{})
}
