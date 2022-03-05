package my

import (
	"errors"
	"log"
)

type Staffs struct {
	Principal
	Teachers
	NonAcademicStaffs
	balance float64
}

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrOverdraftIncurred   = errors.New("overdraft incurred")
)

type Depositable interface {
	Deposit(float64)
}

type Withdrawable interface {
	Withdraw(float64) (float64, error)
}
type BankAccount struct {
	Owner   string
	balance float64
}

func (b BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) (float64, error) {
	if b.balance < amount {
		return 0, ErrInsufficientBalance
	}
	b.balance -= amount
	return b.balance, nil
}

type OverdraftableBankAccount struct {
	BankAccount
	Fee float64
}

func (oba *OverdraftableBankAccount) Withdraw(amount float64) (float64, error) {
	var err error
	if oba.balance < amount {
		oba.balance -= oba.Fee
		err = ErrOverdraftIncurred
	}
	oba.balance -= amount
	return oba.balance, err
}

func Transfer(debtor Withdrawable, creditor Depositable, amount float64) error {
	balance, err := debtor.Withdraw(amount)
	switch err {
	case ErrInsufficientBalance:
		return err
	case ErrOverdraftIncurred:
		log.Printf("debtor incurred overdraft, new balance is %.2f", balance)
	}
	creditor.Deposit(amount)
	return nil
}
