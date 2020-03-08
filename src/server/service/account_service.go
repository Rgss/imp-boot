package service

import (
	"github.com/Rgss/imp-boot/src/server/entity"
	"github.com/gin-gonic/gin"
)

var accountSalt = "!@$ad1"

var AccountServiceFactory = newAccountService()

func newAccountService () *AccountService {
	return &AccountService{}
}

type AccountService struct {
}


//
// login
func (this *AccountService) Login(username string, password string) (*entity.Account, error) {
	var account entity.Account


	return &account, nil
}
