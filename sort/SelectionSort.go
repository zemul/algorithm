package main

/**
SelectionSort
在未排序序列中找到最小（或最大）元素，存放到排序序列的起始位置。
从剩余未排序元素中继续寻找最小（或最大）元素，然后放到已排序序列的末尾。
重复第二步，直到所有元素均排序完毕。
*/

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}
