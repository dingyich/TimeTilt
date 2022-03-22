package coincombination

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
		SuiteName: "coin combination",
	}
}

type suiteData struct {
	coins []int
	gold  int
}

// GetCoinCombination -
// Input coins, total gold
// output all the combinations
func GetCoinCombination(coins []int, gold int) [][]int {
	if len(coins) == 0 || gold < 0 {
		return [][]int{{-1}}
	}

	solution := make([]int, len(coins))
	res := make([][]int, 0)
	res = dfs(coins, res, solution, 0, gold)
	return res
}

func dfs(coins []int, res [][]int, solution []int, index int, gold int) [][]int {
	// base case
	if index == len(coins)-1 {
		r := gold % coins[index]
		if r == 0 {
			solution[index] = gold / coins[index]
			res = append(res, solution)
		}
		return res
	}

	for i := 0; i*coins[index] <= gold; i++ {
		a := deepCopy(solution)
		a[index] = i
		res = dfs(coins, res, a, index+1, gold-i*coins[index])
	}

	return res
}

func deepCopy(solution []int) []int {
	res := make([]int, 0)
	res = append(res, solution...)

	return res
}

func (q *QuzSuite) GetData() suiteData {
	coins := []int{1, 2, 5, 10}
	gold := 5

	return suiteData{
		coins: coins,
		gold:  gold,
	}
}

func (q *QuzSuite) RunSolutionWithData(data suiteData) {
	res := GetCoinCombination(data.coins, data.gold)
	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("Coins: {%v}.\n", data.coins)
	fmt.Printf("Gold: [%v].\n\n", data.gold)
	fmt.Printf("Result: [%v].\n", res)
	fmt.Println("------")
}

func (q *QuzSuite) RunSolution() {
	d := q.GetData()
	res := GetCoinCombination(d.coins, d.gold)
	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("Coins: {%v}.\n", d.coins)
	fmt.Printf("Gold: [%v].\n\n", d.gold)
	fmt.Printf("Result: [%v].\n", res)
	fmt.Println("------")
}
