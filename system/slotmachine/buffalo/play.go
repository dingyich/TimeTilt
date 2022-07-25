package buffalo

import (
	"TimeTilt/system/slotmachine/model"
	"fmt"
	"math"
	"strconv"
	"time"
)

func Run() {
	fmt.Println("Welcome to Buffalo!!!")
	fmt.Println("Good Luck!!!")
	fmt.Println("Please choose credit: 1c 2c 5c 10c 25c 50c 100c. Input number: ")
	var centBase int
	fmt.Scanln(&centBase)
	fmt.Printf("You chose: %vc.\n", centBase)

	fmt.Println("Please choose multiplyer: 1x 2x 3x 4x 5x 6x. Input number: ")
	var multiplyer int
	creditBase := 60
	fmt.Scanln(&multiplyer)
	bet := float64(0.01) * float64(centBase) * float64(multiplyer) * float64(creditBase)
	bet = math.Round(bet*100) / 100
	fmt.Printf("You chose: %vx. Your single bet would be $%.2f.\n", multiplyer, bet)

	bm := NewMachine(model.StartConfig{
		WindowSize:     4,
		CentBase:       centBase,
		BaseCredit:     creditBase,
		SpinMultiplyer: multiplyer,
		Bet:            bet,
	})

	fmt.Println("********************")
	fmt.Printf("* %v c   $ %.2f    *\n", centBase, bet)
	fmt.Println("********************")
	fmt.Printf("Insert bills or tickets. Balance: [$ %.2f].\n", bm.DollarAmount)
	fmt.Println("Press Enter to play... ... ... ...")

	var cmd string
	fmt.Scanln(&cmd)

	for cmd != "exit" {
		switch cmd {
		case "":
			if !bm.EnoughBalance() {
				fmt.Printf("Insert bills or tickets. Balance: [$ %.2f].\n", bm.DollarAmount)
				fmt.Scanln(&cmd)
			} else {
				fmt.Println("********************")
				fmt.Printf("* %v c   $ %.2f    *\n", centBase, bet)
				fmt.Println("********************")

				if bm.FreeGameOn {
					bm.FreeGameStep()
				} else {
					bm.Step()
				}

				if bm.FreeGameOn {
					fmt.Println("Press Enter to start free games...")
					fmt.Scanln(&cmd)
				} else {
					fmt.Println("Press Enter to play... ... ... ...")
					fmt.Scanln(&cmd)
				}
			}

		default:
			s, err := strconv.ParseFloat(cmd, 64)
			if err != nil {
				fmt.Println("Invalid command. Try again!!!")
				time.Sleep(2000 * time.Millisecond)
				fmt.Printf("Insert bills or tickets. Balance: [$ %.2f].\n", bm.DollarAmount)
				fmt.Scanln(&cmd)
			}

			if s != 0 {
				bm.Deposit(s)
				time.Sleep(2000 * time.Millisecond)
				fmt.Printf("Insert bills or tickets. Balance: [$ %.2f].\n", bm.DollarAmount)
				fmt.Println("Press Enter to play... ... ... ...")
				cmd = ""
				fmt.Scanln(&cmd)
			}
		}

	}

	bm.CashOut()
}

func (bm *BuffaloMachine) Step() {
	bm.Shuffle()

	bm.DollarAmount -= bm.Bet
	bm.DollarAmount = math.Round(bm.DollarAmount*100) / 100
	bm.CreditBalance -= int64(bm.CreditBase) * int64(bm.SpinMultiplyer)
	fmt.Printf("Balance: $ %v. Bet: $ %.2f.\n", bm.DollarAmount, bm.Bet)
	fmt.Print("Spinning ")
	for i := 0; i < 3; i++ {
		fmt.Print("-")
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\\")
		time.Sleep(100 * time.Millisecond)
		fmt.Print("|")
		time.Sleep(100 * time.Millisecond)
		fmt.Print("/")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()

	bm.Spin()
	bm.ShowWindow()

	auditResult := bm.Audit()

	if auditResult.ScatterCount > 2 {
		bm.triggerFreeGame(auditResult.ScatterCount)
	}
	bm.ProcessBalance(auditResult)
}

func (bm *BuffaloMachine) triggerFreeGame(scatterCount int) {
	if !bm.FreeGameOn {
		freegames := 8
		if scatterCount == 4 {
			freegames = 15
		} else if scatterCount == 5 {
			freegames = 25
		}
		fmt.Printf("Free Game Won: %d.\n", freegames)
		bm.FreeGames += freegames
		bm.FreeGameOn = true
	} else {
		freegames := 5
		if scatterCount == 3 {
			freegames = 8
		} else if scatterCount == 4 {
			freegames = 15
		} else if scatterCount == 5 {
			freegames = 25
		}
		fmt.Printf("Free Game Won: %d.\n", freegames)
		bm.FreeGames += freegames
	}
}

func (bm *BuffaloMachine) FreeGameStep() {
	bm.FreeGameWinning = 0
	freegameCount := 0
	for bm.FreeGameOn {
		bm.Shuffle()

		freegameCount++
		fmt.Printf("Balance: $ %.2f. \n", bm.DollarAmount)
		fmt.Printf("%d of %d free games spinning ", freegameCount, bm.FreeGames)
		for i := 0; i < 3; i++ {
			fmt.Print("-")
			time.Sleep(100 * time.Millisecond)
			fmt.Print("\\")
			time.Sleep(100 * time.Millisecond)
			fmt.Print("|")
			time.Sleep(100 * time.Millisecond)
			fmt.Print("/")
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()

		bm.Spin()
		bm.ShowWindow()

		auditResult := bm.Audit()
		time.Sleep(1000 * time.Millisecond)

		if auditResult.ScatterCount > 1 {
			bm.triggerFreeGame(auditResult.ScatterCount)
		}
		bm.ProcessBalance(auditResult)

		time.Sleep(1000 * time.Millisecond)

		if freegameCount == bm.FreeGames {
			bm.FreeGameOn = false
			bm.FreeGames = 0
		}
	}

	fmt.Printf("Free game winning: $ %.2f.\n", bm.FreeGameWinning)
	fmt.Println("Free game end.")
	bm.DollarAmount += bm.FreeGameWinning
	bm.DollarAmount = math.Round(bm.DollarAmount*100) / 100
	bm.CreditBalance = int64(bm.DollarAmount * 100 / float64(bm.CentBase))
}

func (bm *BuffaloMachine) ProcessBalance(ar model.AuditResult) {
	var creditEarned int64
	for n, hs := range ar.Stats {
		p := GetPayRate(n)

		num := hs.PayLevel
		index := 0
		if num != 0 {
			index = num - 1
		}
		pays := p[index]

		if pays == 0 {
			continue
		}

		winBase := bm.CreditBase / 60 * bm.SpinMultiplyer

		multi := hs.Multiplyer * ar.Multiplyer
		sleepcount := int(pays * multi * winBase)
		if sleepcount > 2000 {
			sleepcount = 2000
		}
		time.Sleep(time.Duration(sleepcount) * time.Millisecond)
		fmt.Printf("%v [%s]s pays %v x %v = %v.\n", num, n, pays*winBase, multi, pays*multi*winBase)
		creditEarned += int64(pays * multi * winBase)
	}

	if bm.FreeGameOn {
		fmt.Printf("Free game winning: $ %.2f + $ %.2f = $ %.2f.\n", bm.FreeGameWinning, float64(creditEarned)/100, bm.FreeGameWinning+float64(creditEarned)/100)
		bm.FreeGameWinning += float64(creditEarned) / 100
		bm.FreeGameWinning = math.Round(bm.FreeGameWinning*100) / 100
	} else {
		fmt.Printf("Balance: %v + %v = %v.\n", bm.CreditBalance, creditEarned, creditEarned+bm.CreditBalance)
		fmt.Println()
		bm.CreditBalance += creditEarned
		bm.DollarAmount = float64(bm.CreditBalance) * float64(0.01) * float64(bm.CentBase)
		bm.DollarAmount = math.Round(bm.DollarAmount*100) / 100
		fmt.Printf("$ %.2f.\n", bm.DollarAmount)
	}
}
