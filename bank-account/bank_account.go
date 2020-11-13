package account

import "sync"

type Account struct {
	amount int64
	open   bool
	lock   sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, true, sync.Mutex{}}
}

func (a *Account) Balance() (balance int64, ok bool) {
	if !a.open {
		return 0, false
	}
	//a.lock.Lock()
	//defer a.lock.Unlock()
	balance = a.amount
	return balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if !a.open {
		return 0, false
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.amount+amount < 0 {
		return a.amount, false
	}
	a.amount += amount
	newBalance = a.amount
	return newBalance, true
}

func (a *Account) Close() (payout int64, ok bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if !a.open {
		return 0, false
	}
	payout = a.amount
	a.amount = 0
	a.open = false
	return payout, true
}
