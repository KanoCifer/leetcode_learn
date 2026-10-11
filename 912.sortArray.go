package leetcodelearn

import "math/rand"

// 归并排序
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return mergeArr(left, right)
}

func mergeArr(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// 快速排序
func sortArray1(nums []int) []int {
	low, high := 0, len(nums)-1
	quickSort(nums, low, high)

	return nums
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// 优化点
	pivotIndex := low + rand.Intn(high-low+1)
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
