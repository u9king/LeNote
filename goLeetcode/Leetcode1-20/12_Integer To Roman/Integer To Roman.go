package main

//方法一：与以下方法表示结果相同
type romanStruct struct {
	value  int
	symbol string
}

func intToRoman(num int) string {
	romanMap := []romanStruct{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}
	result := ""
	for _, v := range romanMap {
		for num >= v.value {
			result += v.symbol
			num -= v.value
		}
		if num == 0 {
			break
		}
	}
	return result
}

//方法二：使用隐式结构体表达
var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman_v1(num int) (res string) {
	var roman []byte
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}
