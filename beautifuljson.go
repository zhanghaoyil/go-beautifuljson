package go_beautifuljson

import (
	"fmt"
	"io"
	"os"
)

func Add() int {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("读取 %d 字节\n", len(data))
	return 0
}
