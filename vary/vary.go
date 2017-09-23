package vary

import (
	"errors"
	"fmt"
	"os"
)

func GetEnv(key string) (val string, err error) {
	val, ok := os.LookupEnv(key)
	fmt.Println("a", val)
	if ok {
		err = nil
	} else {
		err = errors.New(fmt.Sprintf("系统未设置环境变量%s", key))
	}
	return
}
