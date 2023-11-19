package user

import (
	"context"
	"fmt"

	"github.com/0xfzz/tuwit-backend/domains/user/entity"
	"github.com/0xfzz/tuwit-backend/ent/predicate"
	"github.com/0xfzz/tuwit-backend/lib"
)

type UserRepository struct {
	db lib.Database
}

func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{
		db,
	}
}

func (r UserRepository) FindUserAccount(conditions predicate.UserAccount) (*entity.UserAccountEntity, error) {
	user, err := r.db.DB.UserAccount.Query().Where(conditions).First(context.Background())
	if err != nil {
		return nil, err
	}
	return &entity.UserAccountEntity{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r UserRepository) CreateUser(user entity.UserEntity) error {
	tx, err := r.db.DB.Tx(context.Background())
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	account, err := tx.UserAccount.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(user.Password).Save(context.Background())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("starting a transaction: %w", err)
	}
	_, err = tx.UserProfile.Create().
		SetAccount(account).
		SetDisplayName(user.DisplayName).
		Save(context.Background())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating profile: %w", err)
	}
	tx.Commit()
	return nil
}
