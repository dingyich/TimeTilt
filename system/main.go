package main

import (
	"TimeTilt/system/slotmachine/buffalo"
)

func main() {
	buffalo.Run()

	//bm := buffalo.NewMachine(model.StartConfig{WindowSize: 4})
	//bm.Show()

	// bm := buffalo.MockMachine()
	// bm.ShowWindow()

	// ar := bm.Audit()

	// bm.ProcessBalance(ar)

	// fmt.Println(ar)
}
