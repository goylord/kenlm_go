package main

import "fmt"

func main() {
	model := LoadModel("./test.arpa")
	result := model.Score("测试评分")
	fmt.Println(result)
}
