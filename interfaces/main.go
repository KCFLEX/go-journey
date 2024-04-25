package main

import (
	"example/go-journey/interfaces/wlf"
	"fmt"
)

type BankAccount interface {
	GetBalance() int // 100 = 1 dollar
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	wf := wlf.NewWellsFargo()

	wf.Deposit(200)
	Balance := wf.GetBalance()

	fmt.Println("you deposited ", Balance)

	currentBalance := wf.Withdraw(300)
	fmt.Println("you current balance is now ", currentBalance)

}
