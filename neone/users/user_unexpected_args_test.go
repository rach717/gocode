package users_test

import (
	"github.com/golang/mock/gomock"
	"neone/mocks"
	"neone/users"
	"testing"
)

func TestUser_GoMock_UnexpectedArgs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &users.User{Doer: mockDoer}

	mockDoer.EXPECT().Do(2, "def")

	// Calls mockDoer with (1, "abc")
	testUser.Use()
}

func TestUser_Testify_UnexpectedArgs(t *testing.T) {
	mockDoer := &mocks.Doer{}
	testUser := &users.User{Doer: mockDoer}

	mockDoer.On("Do", 2, "def")

	// Calls mockDoer with (1, "abc")
	testUser.Use()
	mockDoer.AssertExpectations(t)
}
