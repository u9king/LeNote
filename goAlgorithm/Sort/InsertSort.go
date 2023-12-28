package Sort

import (
	"fmt"
)

// InsertSort 插入排序(前小后大)
func InsertSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	} else {
		for i := 1; i < length; i++ {
			backup := arr[i] //备份插入的数据
			j := i - 1
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
		}
	}
}

func main() {
	//输入数据
	list := []int{3,5,10,16,7,32,83,23,54,29,96}

	//输出内容
	InsertSort(list)
	fmt.Println(list)
}

