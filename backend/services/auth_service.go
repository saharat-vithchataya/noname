package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saharat-vithchataya/noname/logs"
	"github.com/saharat-vithchataya/noname/repository"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte("fe819de146e10a641c518f7f0b7890310d98c60ab0daca3e38f3b52cdb6fa855")

// func CreateNewAccessToken()
func createToken(accountID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"account_id": accountID,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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

	// creating token
	accessToken, err := createToken(account.ID)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &LoginResponse{
		AccessToken: accessToken,
	}, nil
}
