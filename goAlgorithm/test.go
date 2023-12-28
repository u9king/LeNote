package main

import (
	"fmt"
	"./Sort"
)

func main() {
	//输入数据
	list := []int{3,5,10,16,7,32,83,23,54,29,96}

	//输出内容
	Sort.InsertSort(list)
	fmt.Println(list)
}
