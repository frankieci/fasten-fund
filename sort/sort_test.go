package sort

import "testing"

/**
 * Created by frankieci on 2021/12/25 8:51 pm
 */

var arr = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

func Test_bubbleSort(t *testing.T) {
	bubbleSort(arr)
	t.Log(arr)
}

func Test_selectSort(t *testing.T) {
	selectSort(arr)
	t.Log(arr)
}

func Test_InsertSort(t *testing.T) {
	insertSort(arr)
	t.Log(arr)
}

func Test_mergeSort(t *testing.T) {
	merged := mergeSort(arr)
	t.Log(merged)
}

func Test_quickSort(t *testing.T) {
	arr = []int{6, 3, 2, 62, 4, 51}
	quickSort(arr)
	t.Log(arr)
}

func Test_shellSort(t *testing.T) {
	shellSort(arr)
	t.Log(arr)
}
