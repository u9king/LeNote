package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 && m > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func removeDuplicates(nums []int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[left] = nums[i]
		}
		if nums[left] != nums[i] {
			left++
			nums[left] = nums[i]
		}
	}
	nums = nums[:left+1]
	return len(nums)
}

func isValidSudoku(board [][]byte) bool {
	//1.数字 1-9 在每一行只能出现一次。
	for i := 0; i < len(board); i++ {
		temp := map[byte]int{'0':1,'1':1,'2':1,'3':1,'4':1,'5':1,'6':1,'7':1,'8':1,'9':1}
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if _,ok := temp[board[i][j]];ok{
				delete(temp, board[i][j])
			} else{
				return false
			}
		}
	}
	
	//2.数字 1-9 在每一列只能出现一次。
	for i := 0; i < len(board); i++ {
		temp := map[byte]int{'0':1,'1':1,'2':1,'3':1,'4':1,'5':1,'6':1,'7':1,'8':1,'9':1}
		for j := 0; j < len(board[i]); j++ {
			if board[j][i] == '.' {
				continue
			}
			if _,ok := temp[board[j][i]];ok{
				delete(temp, board[j][i])
			} else{
				return false
			}
		}
	}
	//3.数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次
	//m := len(board) % 3
	//n := len(board[0]) % 3






	return true
}

func main() {
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	boardboard := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '3'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	isValidSudoku(boardboard)
	fmt.Println(nums1)
}
