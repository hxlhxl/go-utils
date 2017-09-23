package routes

import (
	"fmt"
	"net/http"
	"strings"

	lilyhttp "github.com/hxlhxl/go-utils/http"
)

// 写法很丑，还需要内置的http包，非常的不好看

// type User struct {
// 	lilyhttp.RouteHandler
// }

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	token, ok := r.Header["A"]
	if ok {
		fmt.Println(strings.Split(token[0], ","))
	}
	fmt.Fprintf(w, "user")
}

// post方法,get_user,delete_user
func init() {
	lilyhttp.InitRoute("/api/v1/user", user)
}
