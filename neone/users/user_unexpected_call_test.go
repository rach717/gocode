package users_test

import (
	"github.com/golang/mock/gomock"
	"neone/mocks"
	"neone/users"
	"testing"
)

func TestUser_GoMock_UnexpectedCall(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &users.User{Doer: mockDoer}

	testUser.Use()
}

func TestUser_Testify_UnexpectedCall(t *testing.T) {
	mockDoer := &mocks.MockDoer{}
	testUser := &users.User{Doer: mockDoer}

	testUser.Use()

}

