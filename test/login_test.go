package test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/0xfzz/tuwit-backend/domains/auth/dto"
	"github.com/0xfzz/tuwit-backend/domains/user"
	"github.com/0xfzz/tuwit-backend/domains/user/entity"
	"github.com/0xfzz/tuwit-backend/test/mock"
	"github.com/0xfzz/tuwit-backend/utils/hash"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		db := mock.NewMockDatabase(t)
		r := mock.RouterAuthController(t, db)
		var entiti = dto.RegisterRequest{
			Username:        "yanto",
			Email:           "yanto2@test.com",
			DisplayName:     "yantoh",
			Password:        "testtesttesttest",
			ConfirmPassword: "testtesttesttest",
		}
		require.NoError(t, db.DB.Schema.Create(context.Background()))
		userRepo := user.NewUserRepository(db)
		var err = userRepo.CreateUser(entity.UserEntity{
			Username:    entiti.Username,
			Email:       entiti.Email,
			DisplayName: entiti.DisplayName,
			Password:    hash.Hash(entiti.Password),
		})
		if err != nil {
			t.Errorf("got an error : %s", err)
		}
		w := httptest.NewRecorder()
		bodyString := fmt.Sprintf(`{"identifier":"%s", "password":"%s"}`, entiti.Email, entiti.Password)
		req, err := http.NewRequest("POST", "/api/auth/login", bytes.NewReader([]byte(bodyString)))
		r.ServeHTTP(w, req)
		result := w.Result()
		resp, _ := io.ReadAll(result.Body)
		t.Log(string(resp))
		if result.StatusCode != 200 {
			t.Errorf("Expected Status Code 200 but got %d", w.Code)
		}
		if err != nil {
			t.Errorf("got an error : %s", err)
		}
	})
}
