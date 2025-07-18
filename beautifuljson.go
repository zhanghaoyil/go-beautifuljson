package go_beautifuljson

import (
	"fmt"
	"os"
)

func Add() int {
	content, err := os.ReadFile("/etc/hosts")
	if err != nil {
		fmt.Println("读取文件出错:", err)
		return -1
	}
	fmt.Println(string(content))
	return 0
}
