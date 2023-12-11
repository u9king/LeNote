package main

//romanToInt 使用字典更好理解一些，switch方法执行效率更高可以尝试
func romanToInt(s string) int {
	romanMap := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400,
		"C": 100, "XC": 90, "L": 50, "XL": 40,
		"X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	result := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if i < n-1 {
			double := s[i : i+2]
			if cur, ok := romanMap[double]; ok {
				result += cur
				i++
				continue
			}
		}
		single := s[i : i+1]
		result += romanMap[single]
	}
	return result
}

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt_v1(s string) (ans int) {
	n := len(s)
	for i := 0; i < n; i++ {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans

}
