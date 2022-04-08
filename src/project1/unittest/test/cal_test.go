package test

import (
	"fmt"
	"testing"
)

//单元测试，必须以Test开头
//单元测试文件必须xxx_test.go的方式命名
//测试方法：go run -v
func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 10 {
		t.Fatalf("AddUpper(10)执行错误，期望值=%v，实际值=%v\n", 10, res)
	}
	t.Logf("AddUpper(10)执行正确")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用...")
}
