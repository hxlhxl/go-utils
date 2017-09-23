package file

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func IsExist(filePath string) bool {
	_,err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}

func EchoFileName(filename string) {
	fmt.Println(filename)
}