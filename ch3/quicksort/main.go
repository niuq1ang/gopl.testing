package main

import "fmt"

func main() {
	arr := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22}
	quickSort(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}

// 快速排序使用分治的思想，通过一趟排序将待排序列分割成两部分，其中一部分记录的关键字均比另一部分记录的关键字小。
// 之后分别对这两部分记录继续进行排序，以达到整个序列有序的目的
func quickSort(array []int) {
	if len(array) <= 1 {
		return
	}

	/*
	* 快速排序步骤：
	* 选择基准：在待排序列中，按照某种方式挑出一个元素，作为 "基准"（pivot）
	* 分割操作：以该基准在序列中的实际位置，把序列分成两个子序列。此时，在基准左边的元素都比该基准小，在基准右边的元素都比基准大
	* 递归地对两个序列进行快速排序，直到序列为空或者只有一个元素。
	 */
	pivot, left, right := array[0], 0, len(array)-1

	for left < right {
		for left < right && array[right] >= pivot {
			right--
		}
		array[left] = array[right]

		for left < right && array[left] < pivot {
			left++
		}
		array[right] = array[left]
	}
	array[left] = pivot
	quickSort(array[:left])
	quickSort(array[left+1:])
}
