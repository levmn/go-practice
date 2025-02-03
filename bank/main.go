package main

import "fmt"

type CheckingAccount struct {
	owner         string
	agency        int
	accountNumber int
	balance       float64
}

func (c *CheckingAccount) Withdraw(withdrawalValue float64) (string, float64) {
	allowedWithdrawal := withdrawalValue > 0 && withdrawalValue <= c.balance

	if allowedWithdrawal {
		c.balance -= withdrawalValue
		return "Withdrawal successful! Current account balance:", c.balance
	} else {
		return "Insufficient funds! Current account balance:", c.balance
	}
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successful! Current account balance:", c.balance
	} else {
		return "Invalid deposit value! Current account balance:", c.balance
	}
}

func (c *CheckingAccount) Transfer(transferValue float64, destinationAccount *CheckingAccount) bool {
	if transferValue < c.balance {
		c.balance -= transferValue
		destinationAccount.balance += transferValue
		return true
	} else {
		return false
	}
}

func main() {
	joaoAccount := CheckingAccount{owner: "João", agency: 589, accountNumber: 123456, balance: 125.50}
	mariaAccount := CheckingAccount{owner: "Maria", agency: 589, accountNumber: 654321, balance: 1000}

	fmt.Println(joaoAccount)

	fmt.Println(joaoAccount.Withdraw(100))

	fmt.Println(joaoAccount.Deposit(500))

	fmt.Println(joaoAccount.Transfer(50, &mariaAccount))

	fmt.Println(mariaAccount)
}
