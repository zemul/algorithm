package main

// InsertionSort 类似扑克牌的手牌排序方法，起始时将头部元素作为有序数组，剩下部分作为无序数组
// 从无序数组中依次取值作为候选值(candidate)，然后倒叙遍历有序数组，每当发现候选值更小那么交换一次位置，否则直接break
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr)-1; i++ {
		j := i - 1
		// 将选中的元素与已排序的元素从后向前比较，找到合适的插入位置
		for j >= 0 && arr[j] > arr[i] {
			arr[j+1] = arr[j] // 元素后移
			j--
		}
		arr[j+1] = arr[i]
	}
	return arr
}
