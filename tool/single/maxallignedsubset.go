package single

import "sort"

func GetMaxAllignedSubset(arr []int, m int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if m <= 0 {
		return []int{}
	}

	sortArray(arr, true)
	res := make(map[int][]int)
	bias := arr[0]
	maxLength := 1
	current := 0
	for _, i := range arr {
		a := (i - bias) % m
		list := res[a]
		list = append(list, i)
		res[a] = list

		if len(list) > maxLength {
			maxLength = len(list)
			current = a
		}
	}

	return res[current]
}

func sortArray(arr []int, asc bool) {
	sort.Slice(arr, func(i, j int) bool {
		mi := arr[i]
		mj := arr[j]

		if asc {
			return mi < mj
		}

		return mj < mi
	})
}
