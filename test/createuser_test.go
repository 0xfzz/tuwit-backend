package test

import (
	"context"
	"testing"

	"github.com/0xfzz/tuwit-backend/domains/user"
	"github.com/0xfzz/tuwit-backend/domains/user/entity"
	"github.com/0xfzz/tuwit-backend/ent/useraccount"
	"github.com/0xfzz/tuwit-backend/test/mock"
	"github.com/0xfzz/tuwit-backend/utils/hash"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	t.Run("create_user", func(t *testing.T) {
		db := mock.NewMockDatabase(t)
		require.NoError(t, db.DB.Schema.Create(context.Background()))
		userRepo := user.NewUserRepository(db)
		password := hash.Hash("PAsswordKeren")
		var err = userRepo.CreateUser(entity.UserEntity{
			Username:    "foo",
			DisplayName: "Phat",
			Email:       "test@test.com",
			Password:    password,
		})
		if err != nil {
			t.Errorf("Got an error : %s", err)
		}
		user, err := userRepo.FindUserAccount(useraccount.Email("test@test.com"))
		if err != nil {
			t.Errorf("Got an error : %s", err)
		}
		expectedUserAccount := entity.UserAccountEntity{
			Username: "foo",
			Email:    "test@test.com",
		}
		returnUserAccount := entity.UserAccountEntity{
			Username: user.Username,
			Email:    user.Email,
		}

		if !cmp.Equal(returnUserAccount, expectedUserAccount, cmpopts.IgnoreFields(entity.UserAccountEntity{}, "Password")) {
			t.Errorf("Expected %+v got %+v", expectedUserAccount, returnUserAccount)
		}

	})

}
