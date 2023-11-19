package auth

import (
	"github.com/0xfzz/tuwit-backend/domains/auth/dto"
	"github.com/0xfzz/tuwit-backend/domains/user"
	"github.com/0xfzz/tuwit-backend/domains/user/entity"
	"github.com/0xfzz/tuwit-backend/ent/useraccount"
	"github.com/0xfzz/tuwit-backend/utils/hash"
	"github.com/0xfzz/tuwit-backend/utils/token"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepo user.UserRepository
}

func NewAuthService(userRepo user.UserRepository) AuthService {
	return AuthService{
		userRepo,
	}
}
func (s AuthService) CheckUsernameAvailability(username string) bool {
	_, err := s.userRepo.FindUserAccount(
		useraccount.Username(username),
	)

	return err == nil

}
func (s AuthService) RegisterUser(userReq dto.RegisterRequest) error {
	hashedPassword := hash.Hash(userReq.Password)
	return s.userRepo.CreateUser(entity.UserEntity{
		Username:    userReq.Username,
		Email:       userReq.Email,
		Password:    hashedPassword,
		DisplayName: userReq.DisplayName,
	})
}
func (s AuthService) LoginUser(identifier string, password string) (*string, error) {
	user, err := s.userRepo.FindUserAccount(
		useraccount.Or(
			useraccount.Email(identifier),
			useraccount.Username(identifier),
		),
	)
	if err != nil || !hash.Compare(user.Password, []byte(password)) {
		return nil, err
	}
	token := token.Sign(jwt.MapClaims{
		"user_id": user.Id,
	})
	return &token, nil
}
