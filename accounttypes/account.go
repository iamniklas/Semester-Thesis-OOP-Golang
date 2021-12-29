package accounttypes

import (
	"container/list"
	"time"
)

type Account struct {
	firstName             string
	lastName              string
	accountIdentifier     string
	pin                   string
	lastTimeLoggedIn      time.Time
	accountBalance        float32
	transactionHistory    list.List
	transactionMultiplier float32
}

type BasicAccount struct {
	Account
}

type StandardAccount struct {
	Account
}

type SuperPremiumAccount struct {
	Account
}
