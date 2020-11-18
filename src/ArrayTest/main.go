package main

/*
数组，是同一种数据类型的固定长度的序列
定义： var a [len] int 
var a [5] int // 数组长度必须是常量，

*/
import (
	"fmt"
)
var arr0 [5]int = [5]int{1,2,3}
var arr1 = [5]int{1,2,3,4,5}
var arr2 = [...]int{1,2,3,4,5,6}
var str = [5]string{3: "hello world ",4: "tom"}
func main() {
	a := [3]int{1, 2}// 未初始化元素值为 0
	b := [...]int{1, 2, 3, 4} // 通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素
	d := [...]struct {
		name string
		age uint8
	}{
		{"user1", 10}, //可省略元素类型
		{"user2", 20},  // 最后一行逗号
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}

