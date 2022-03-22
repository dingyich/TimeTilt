package forbiddentrioswap

import "fmt"

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
		SuiteName: "forbidden trio swap",
	}
}

type suiteData struct {
	nums []int
}

func ForbiddenTrioSwap(nums []int) int {
	res := 0

	if len(nums) < 3 {
		return res
	}

	i := 0
	for i < len(nums) {
		next := i + 1
		for next < len(nums) && nums[next] == nums[i] {
			next++
		}

		res += (next - i) / 3
		i = next
	}

	return res
}

// ForbiddenTrioSwap2 can only swap elements
func ForbiddenTrioSwap2(nums []int) int {
	if !ValidateFTS2(nums) {
		return -1
	}

	if len(nums) < 3 {
		return 0
	}

	res := 0
	return res
}

// ValidateFTS2 desc
// nums has at most 2 different numbers
// using swap to make sure no such case: a,a,a or b,b,b
func ValidateFTS2(nums []int) bool {
	if len(nums) < 3 {
		return true
	}

	n := nums[0]
	a := 0
	b := 0
	for _, i := range nums {
		if i == n {
			a++
		} else {
			b++
		}
	}

	if a < b {
		a, b = b, a
	}

	if (a-1)/2 > b {
		return false
	}
	return true
}

func (q *QuzSuite) GetData() suiteData {
	nums := []int{4, 9, 9, 9, 9, 9, 9, 4, 9, 9, 4}

	return suiteData{
		nums: nums,
	}
}

func (q *QuzSuite) RunSolutionWithData(data suiteData) {
	res := ValidateFTS2(data.nums)
	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("nums: {%v}.\n\n", data.nums)
	fmt.Printf("Result: [%v].\n", res)
	fmt.Println("------")
}

func (q *QuzSuite) RunSolution() {
	d := q.GetData()
	res := ValidateFTS2(d.nums)
	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("nums: {%v}.\n\n", d.nums)
	fmt.Printf("Result: [%v].\n", res)
	fmt.Println("------")
}
