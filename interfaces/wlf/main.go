package wlf

import "log"

type WellsFargo struct {
	balance int
}

func NewWellsFargo() *WellsFargo { // constructor
	return &WellsFargo{
		balance: 0,
	}
}

func (w *WellsFargo) GetBalance() int {
	return w.balance
}

func (w *WellsFargo) Deposit(amount int) {
	w.balance += amount
}

func (w *WellsFargo) Withdraw(amount int) int {
	newBalance := w.balance - amount
	if newBalance < 0 {
		log.Println("insufficient funds")
	}
	return newBalance
}
