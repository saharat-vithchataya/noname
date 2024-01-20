package services

import (
	"github.com/saharat-vithchataya/noname/logs"
	"github.com/saharat-vithchataya/noname/repository"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	accRepo repository.AccountRepository
}

func NewAuthService(accRepo repository.AccountRepository) AuthService {
	return authService{accRepo: accRepo}
}

func (srv authService) Login(userCredentials LoginScheme) (*LoginResponse, error) {

	account, err := srv.accRepo.GetAccountByEmail(userCredentials.Email)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(userCredentials.Password)); err != nil {
		logs.Error(err)
		return nil, err
	}

	return &LoginResponse{
		AccessToken: "wewe",
	}, nil
}
