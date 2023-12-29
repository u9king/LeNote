package main

import (
	"fmt"
)

type SString struct {
	data   string
	length int
}

// SubString 返回一个子串
func (s SString) SubString(start int, length int) SString {
	if start < 0 || start >= s.length || length <= 0 {
		return SString{"", 0}
	}

	end := start + length
	if end > s.length {
		end = s.length
	}

	return SString{
		data:   s.data[start:end],
		length: end - start,
	}
}

func main() {
	str := SString{
		data:   "Hello, World!",
		length: 13,
	}

	substr := str.SubString(7, 5)
	fmt.Printf("Substring: %s\n", substr.data)
	fmt.Printf("Substring Length: %d\n", substr.length)
}
