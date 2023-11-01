package numberguess

import (
	"fmt"
	"math/rand"
	"time"
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
		SuiteName: "number guess",
	}
}

type suiteData struct {
	NumOfBits           int
	PositionRightSymbol string
	NumberRightSymbol   string
}

func (q *QuzSuite) GetData() suiteData {
	return suiteData{
		NumOfBits:           4,
		PositionRightSymbol: POSITIONCORRECT,
		NumberRightSymbol:   NUMBERCORRECT,
	}
}

func (q *QuzSuite) RunSolutionWithData(data suiteData) {

	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("Digits: [%v].\n\n", data.NumOfBits)

	targetNumber := q.generateRandomNumber(data.NumOfBits)
	q.gamePlay(data, targetNumber)
}

func (q *QuzSuite) RunSolution() {
	data := q.GetData()

	fmt.Printf("Running quiz suite: [%s].\n", q.SuiteName)
	fmt.Printf("Digits: [%v].\n\n", data.NumOfBits)

	targetNumber := q.generateRandomNumber(data.NumOfBits)
	q.gamePlay(data, targetNumber)
}

func (q *QuzSuite) generateRandomNumber(bits int) []int8 {
	res := make([]int8, 0, bits)
	rand.Seed(time.Now().UnixNano())
	tens := 1
	for i := 0; i < bits; i++ {
		tens = tens * 10
	}
	n := rand.Intn(tens)

	tens = tens / 10

	for i := 0; i < bits; i++ {
		d := n / tens

		res = append(res, int8(d+48))

		n = n - tens*d
		tens = tens / 10
	}

	return res
}

func (q *QuzSuite) checkNumber(postionSymbol, numberSymbol string, tagetNumber, guessNumber []int8) (string, bool) {
	res := ""

	position := 0
	number := 0

	helperMap := make(map[int8]int)
	helperTarget := make([]int8, 0, len(tagetNumber))
	helperGuess := make([]int8, 0, len(guessNumber))

	for i, n := range guessNumber {
		if n == tagetNumber[i] {
			position++
		} else {
			helperTarget = append(helperTarget, tagetNumber[i])
			helperGuess = append(helperGuess, n)
		}
	}

	for _, i := range helperTarget {
		helperMap[i]++
	}

	for _, n := range helperGuess {
		if found := helperMap[n]; found > 0 {
			number++
			helperMap[n]--
		}
	}

	for i := 0; i < position; i++ {
		res += postionSymbol + " "
	}

	if position == len(tagetNumber) {
		return res, true
	}

	for i := 0; i < number; i++ {
		res += numberSymbol + " "
	}

	return res, false
}

func (q *QuzSuite) gamePlay(data suiteData, targetNumber []int8) {
	numberOfTry := 0

	fmt.Printf("Enter Number to Guess...\n")

	var cmd string
	fmt.Scanln(&cmd)

	for cmd != "exit" && cmd != "quit" {
		guessNumber, valid := q.buildGuessNumber(cmd)
		if !valid {
			fmt.Println("Invalid Input ... ")
			cmd = ""
			fmt.Scanln(&cmd)
		}
		if len(guessNumber) != len(targetNumber) {
			fmt.Println("Invalid Input ... ")
			cmd = ""
			fmt.Scanln(&cmd)
		}

		checkResult, success := q.checkNumber(data.PositionRightSymbol, data.NumberRightSymbol, targetNumber, guessNumber)

		fmt.Printf("%s   %s   \n", cmd, checkResult)
		//fmt.Println(targetNumber)

		numberOfTry++
		if success {
			fmt.Printf("Success Guess. Number of tries: [%d].\n", numberOfTry)
			break
		}
		cmd = ""
		fmt.Scanln(&cmd)
	}
}

func (q *QuzSuite) buildGuessNumber(guess string) ([]int8, bool) {
	res := make([]int8, 0, len(guess))
	for _, c := range guess {
		n := int8(c)

		if n < 48 || n > 57 {
			return res, false
		}

		res = append(res, n)
	}

	return res, true
}
