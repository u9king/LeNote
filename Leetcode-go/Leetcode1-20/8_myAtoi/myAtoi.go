package main

func myAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1
	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}

	for ; i < n; i++ {
		if s[i] < 48 || s[i] > 57 {
			break
		}
		result = result*10 + int(s[i]-'0')
		if sign*result < MinInt32 {
			return MinInt32
		}
		if sign*result > MaxInt32 {
			return MaxInt32
		}
	}
	return sign * result
}

func myAtoi_waitTest(s string) int {
	//错在哪里？

	result := 0
	sign := 1        //符号位
	hasSign := false //出现例子：“+-12”
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1
	for _, value := range s {
		if value == ' ' || value == '0' && result == 0{ //出现空或者左边的0就跳过
			continue
		}
		if value == '-' && hasSign == false { //符号保存下来
			sign = -1
			hasSign = true
			continue
		} else if value == '+' && hasSign == false { //出现例子："+1"
			hasSign = true
			continue
		}

		if value >= '0' && value <= '9' {
			result = result*10 + int(value-'0') //进位并存储新数用ASCII码表示
			if result*sign > MaxInt32 {         //判断是否大于int32最大数
				return MaxInt32
			}
			if result*sign < MinInt32 { //判断是否小于int32最小数
				return MinInt32
			}
		} else {
			break
		}
	}
	return result * sign
}