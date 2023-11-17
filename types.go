package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct{
	Number 		int64 `json:"number"`
	Password 	string `json:"password"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount  int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        			int 		`json:"id"`
	FirstName 			string 		`json:"firstName"`
	LastName  			string 		`json:"lastName"`
	CreatedAt 			time.Time 	`json:"createdAt"`
	Number	  			int64		`json:"number"`
	EncryptedPassword 	string 		`json:"-"`
	Balance	  			int64		`json:"balance"`
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName: 			firstName,
		LastName: 			lastName,
		Number:				int64(rand.Intn(1000000)), 
		EncryptedPassword: 	string(encpw),
		CreatedAt: 			time.Now().UTC(),
	}, nil
}