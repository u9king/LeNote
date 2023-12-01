package main

import "fmt"

func convert(s string, numRows int) string {
	rows := make([]string, numRows)
	rowNum := 0
	step := 1 //带方向
	for _, v := range s {
		rows[rowNum] += string(v)	//同行字符串添加

		if numRows-1 == 0 {	//处理只有一行的问题
			step = 0
		} else if rowNum == numRows-1 {
			step = -1
		} else if rowNum == 0 {
			step = 1
		}
		rowNum += step
	}
	//复原字符串
	var res string
	for _, row := range rows {
		res += row
	}
	return res
}

func main() {
	//输入数据
	s := "AB"

	//输出内容
	fmt.Println(convert(s, 1))
}
