package test

import (
	"fmt"

	"github.com/goylord/kenlm_go"
)

func main() {
	model := kenlm_go.LoadModel("./test.arpa")
	result := model.Score("测试评分")
	fmt.Println(result)
}
