package single

func MinMoves2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	nums = quickSort(nums, 0, len(nums)-1)
	i := len(nums) / 2
	mid := nums[i]

	sum := 0
	for _, n := range nums {
		if n < mid {
			sum += mid - n
		} else {
			sum += n - mid
		}
	}

	return sum
}

func quickSort(nums []int, left int, right int) []int {
	// base case
	if left >= right {
		return nums
	}

	pivot := baseQuickSort(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)

	return nums
}

func baseQuickSort(nums []int, left int, right int) int {
	pivot := nums[right]

	i := left
	j := right - 1
	for i <= j {
		if nums[i] < pivot {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}
