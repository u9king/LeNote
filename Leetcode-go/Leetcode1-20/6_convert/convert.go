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
