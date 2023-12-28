package main

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}
	T := 2*numRows - 2
	c := (n + T - 1) / T * (numRows - 1)
	mat := make([][]byte, numRows)
	for i := range mat {
		mat[i] = make([]byte, c)
	}
	x, y := 0, 0
	for i := range s {
		mat[x][y] = byte(s[i])
		if i%T < numRows-1 {
			x++
		} else {
			x--
			y++
		}
	}
	res := []byte{}
	for _, row := range mat {
		for _, ch := range row {
			if ch != 0 {
				res = append(res, ch)
			}
		}
	}
	return string(res)
}

//convert_v1 Z字变换，理解step方向和数学原理用一维数组实现
func convert_v1(s string, numRows int) string {
	rows := make([]string, numRows)
	rowNum := 0	//行编号
	step := 1 //带方向的步长
	for _, v := range s {
		rows[rowNum] += string(v) //同行字符串添加
		if numRows-1 == 0 { //处理只有一行的问题
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
