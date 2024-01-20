package services

import (
	"database/sql"
	"time"

	"github.com/saharat-vithchataya/noname/logs"
	"github.com/saharat-vithchataya/noname/repository"
	"golang.org/x/crypto/bcrypt"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (srv accountService) OpenNewAccount(newAcc CreateNewAccountScheme) (*AccountResponseScheme, error) {

	// checking email is already exists, isn't it
	_, err := srv.accRepo.GetAccountByEmail(newAcc.Email)
	if err != sql.ErrNoRows {
		logs.Error(ErrEmailAlreadyExists)
		return nil, ErrEmailAlreadyExists
	}

	// hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAcc.Password), 10)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// creating an account
	user, err := srv.accRepo.Create(repository.Account{
		Email:          newAcc.Email,
		HashedPassword: string(hashedPassword),
		IsVerified:     false,
		OpeningDate:    time.Now().Format("2006-1-2 15:04:05"),
	})
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &AccountResponseScheme{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (srv accountService) GetAccount(accountID int) (*AccountResponseScheme, error) {
	account, err := srv.accRepo.GetAccount(accountID)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &AccountResponseScheme{
		ID:    accountID,
		Email: account.Email,
	}, nil
}
