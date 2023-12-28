package Sort

import (
	"fmt"
)

//InsertSort 插入排序(前小后大)
//重点：内圈循坏条件以及移动方法
func InsertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ { 		   //从第二个位置开始排
		backup := arr[i]                       //备份待插入的数据
		j := i - 1                             //从前一个已经排过序的开始找
		for ; j >= 0 && backup < arr[j]; j-- { //从后往前检索插入位置
			arr[j+1] = arr[j] 		//一个一个后移
		}
		arr[j+1] = backup
	}
}

//InsertSortDummy 带哨兵版本
func InsertSortDummy(arr []int) {
	if len(arr) <= 1 {
		return
	}
	newArr := append([]int{0}, arr...) //增加哨兵位
	for i := 2; i < len(newArr); i++ {
		if newArr[i] < newArr[i-1] {
			newArr[0] = newArr[i]
			j := i - 1
			for ; newArr[0] < newArr[j]; j-- { //从后往前检索插入位置
				newArr[j+1] = newArr[j] //整体后移
			}
			newArr[j+1] = newArr[0] //备份待插入的数据
		}
	}
	copy(arr, newArr[1:]) //深拷贝
}

func main() {
	//输入数据
	list := []int{3, 5, 10, 16, 7, 32, 83, 23, 54, 29, 96}

	//输出内容
	InsertSort(list)
	fmt.Println(list)
}
