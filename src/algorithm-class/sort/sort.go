package main

import "fmt"

func main() {
	arr := []int{2, 15, 78, 9, 13, 58}
	bubbleSort(arr)

	fmt.Println("bubbleSort: ", arr)

}

// bubbleSort 冒泡排序，依次比较和排序。把大的或小的一步步沉淀到后面去
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
