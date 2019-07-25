package ken_logging

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

func checkPath(filePath string) {
	if !isDirExists(filePath) {
		_ = os.MkdirAll(filepath.Dir(filePath), 0666)
	}
}

func GetCurrentDirectory() string {
	// 返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
