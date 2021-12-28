package sort

/**
 * Created by frankieci on 2021/12/25 8:50 pm
 */

func bubbleSort(list []int) {
	length := len(list)
	swapped := false

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}

func selectSort(list []int) {
	length := len(list)
	// 进行 N-1 轮迭代
	for i := 0; i < length-1; i++ {
		// 每次从第 i 位开始，找到最小的元素
		min := list[i] // 最小数
		minIndex := i  // 最小数的下标
		for j := i + 1; j < length; j++ {
			if list[j] < min {
				// 如果找到的数比上次的还小，那么最小的数变为它
				min = list[j]
				minIndex = j
			}
		}

		// 这一轮找到的最小数的下标不等于最开始的下标，交换元素
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func insertSort(list []int) {
	length := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= length-1; i++ {
		spec := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		if spec < list[j] { // 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理

			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && spec < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = spec // 结束了，待排序的数插入空位
		}
	}
}

func mergeSort(list []int) []int {
	length := len(list)
	if length == 1 {
		return list //最后切割只剩下一个元素
	}
	m := length / 2
	left := mergeSort(list[:m])
	right := mergeSort(list[m:])
	return merge(left, right)
}

func merge(l []int, r []int) []int { //把两个有序切片合并成一个有序切片

	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)

	lIndex, rIndex := 0, 0 //两个切片的下标，插入一个数据，下标加一
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}

func quickSort(list []int) {
	length := len(list)
	if length < 2 {
		return
	}
	head, trip := 0, length-1
	for head < trip { //list[head]就是我们的标尺，
		if list[head+1] > list[head] { //标尺元素遇到大于它的，就把这个元素丢到最右边trip
			list[head+1], list[trip] = list[trip], list[head+1]
			trip--
		} else if list[head+1] < list[head] { //标尺元素遇到小于它的，就换位置，标尺右移动一位。
			list[head], list[head+1] = list[head+1], list[head]
			head++
		} else { //相等不用交换
			head++
		}
	}
	quickSort(list[:head])
	quickSort(list[head+1:])
}

func shellSort(list []int) {
	length := len(list)
	h := 1
	for h < length/3 { //寻找合适的间隔h
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < length; i++ { //将数组变为间隔h个元素有序
			for j := i; j >= h && list[j] < list[j-h]; j -= h { //间隔h插入排序
				swap(list, j, j-h)
			}
		}
		h /= 3
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

/*
func binarySearch(target int, list []int) int {
	if len(list) < 0 {
		return -1
	}
	left, right := 0, len(list)-1
	for left <= right {
		mid := (left + right) / 2
		if list[mid] == target {
			return mid
		}
		if list[mid] > target {
			right = mid - 1
		}
		if list[mid] < target {
			left = mid + 1
		}
	}

	return -1
}*/
