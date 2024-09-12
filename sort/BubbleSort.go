package main

import "fmt"

/**
最差时间复杂度皆为O(n²)，空间复杂度皆为O(1)。
三种基础排序算法的优劣？

冒泡排序总是需要多次的比较，替换次数和数组的有序度强相关，数组越有序需要的替换次数越少。
选择排序总是需要多次的比较，替换次数是固定的。
插入排序的比较和替换次数与数组的有序度相关，在数组有序的场景下，比较和替换的次数都会极低。
因此插入排序在实践中一般会比冒泡和选择排序更快
*/

func main() {
	// 示例数组
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Original array:", arr)

	BubbleSort2(arr)
	fmt.Println("Sorted array:", arr)
}

// BubbleSort 挨个排序相邻元素，最终末尾就是最大值，把去除末尾元素的数组作为新的未排序数组，重复此行为直到未排序数组为空
func BubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// BubbleSort1 优化1 没有交换跳出排序
func BubbleSort1(array []int) []int {
	for i := 0; i < len(array); i++ {
		var change bool
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				change = true
			}
		}
		if !change {
			break
		}
	}
	return array
}

// BubbleSort2 优化2,可以记下最后一次交换的位置，后边没有交换，必然是有序的，然后下一次排序从第一个比较到上次记录的位置结束即可。
func BubbleSort2(array []int) []int {
	lastExchange := len(array) - 1
	for i := 0; i < len(array); i++ {
		newLastExchange := 0 // 每次循环重置新交换的位置
		for j := 0; j < lastExchange; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				newLastExchange = j
			}
		}
		lastExchange = newLastExchange
	}
	return array
}
