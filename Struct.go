package main

import "fmt"

func main() {
	myAccount := Account{
		name: "nhan",
		balance: 100,
	}

	myAccount.bonusBalance(20)
	myAccount.bonusBalance(30)
	myAccount.bonusBalance(20)
	myAccount.bonusBalance(30)

	fmt.Println(myAccount.getBalance())

}


type Account struct {
	name string
	balance int
}

func (a Account)getBalance ()int {
	return a.balance
}

func (a *Account) bonusBalance (amount int ) int {
	a.balance += amount
	return a.balance
}