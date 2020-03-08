package service

import (
	"github.com/Rgss/imp-boot/src/common/app"
	"github.com/Rgss/imp-boot/src/server/entity"
	"github.com/Rgss/imp-boot/src/server/repository"
)

// token 有限期
const tokenExpire = 60 * 60 * 24 * 30

// 是否允许多设备登陆
const allowedMutil = false

// token service
var TokenServiceFactory = newTokenService()

type TokenService struct {
}

// factory
func newTokenService() *TokenService {
	return &TokenService{}
}

// get one Token
func (this *TokenService) GetToken(id int) *entity.Token {
	token := &entity.Token{}
	_, err := repository.TokenRepository.Find(id, token)
	if err != nil {
		app.Logger().Error("TokenService.GetToken err: %v", err.Error())
	}

	return token
}