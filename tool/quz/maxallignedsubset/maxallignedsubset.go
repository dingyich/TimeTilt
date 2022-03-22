package maxallignedsubset

import (
	"fmt"
	"sort"
)

type IQuzSuite interface {
	GetData() suiteData
	RunSolution()
	RunSolutionWithData(data suiteData)
}

type QuzSuite struct {
	SuiteName string
}

func NewQuzSuite() IQuzSuite {
	return &QuzSuite{
		SuiteName: "max alligned subset",
	}
}

type suiteData struct {
	nums         []int
	allignnumber int
}

// GetMaxAllignedSubset -
// 1,4,7,10 are alligned set because the distance between each point are divisible by 3
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

func (q *QuzSuite) GetData() suiteData {
	nums := []int{-1, 10, 99, 1, 2, 53, 17, 40, 5, -8, 10}
	m := 3

	return suiteData{
		nums:         nums,
		allignnumber: m,
	}
}

func (q *QuzSuite) RunSolutionWithData(data suiteData) {
	res := GetMaxAllignedSubset(data.nums, data.allignnumber)
	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("nums: {%v}.\n", data.nums)
	fmt.Printf("alligned number: {%v}.\n\n", data.allignnumber)
	fmt.Printf("Result: [%v].\n", res)
	fmt.Println("------")
}

func (q *QuzSuite) RunSolution() {
	d := q.GetData()
	res := GetMaxAllignedSubset(d.nums, d.allignnumber)
	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("nums: {%v}.\n", d.nums)
	fmt.Printf("alligned number: {%v}.\n\n", d.allignnumber)
	fmt.Printf("Result: [%v].\n", res)
	fmt.Println("------")
}
