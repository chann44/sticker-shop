package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chann44/go-shop/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userMocStore := &MocUserStore{}
	handler := NewHandler(userMocStore)

	t.Run("should fail if the user ID is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user/abc", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/user/{userID}", handler.handleGetUser).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

}

type MocUserStore struct {
}

func (ms *MocUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (ms *MocUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (ms *MocUserStore) CreateUser(user types.User) error {
	return nil
}
