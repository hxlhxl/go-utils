package path

import (
	"log"
	"os"
)

// IsExist 判断文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// GetExewd 返回当前运行程序的 路径
func GetExewd() string {
	var ExeRoot string
	var err error
	ExeRoot, err = os.Executable()
	if err != nil {
		log.Fatalln("GetExewd fail:", err)
	}
	return ExeRoot
}

// Getwd 返回当前程序的工作路径
func Getwd() string {
	var Root string
	var err error
	Root, err = os.Getwd()
	if err != nil {
		log.Fatalln("Getwd fail:", err)
	}
	return Root
}
