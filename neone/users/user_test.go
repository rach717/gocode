package users_test

import (
	"github.com/golang/mock/gomock"
	"neone/mocks"
	"neone/users"
	"testing"
)

func TestUserWithGoMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &users.User{Doer:mockDoer}

	// Expect Do to be called once with 1 and "abc" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().Do(1, "abc").Return(nil).Times(1)

	testUser.Use()
	mockCtrl.Finish()
}

func TestUserWithTestifyMock(t *testing.T) {
	mockDoer := &mocks.Doer{}

	testUser := &users.User{Doer:mockDoer}

	// Expect Do to be called once with 1 and "abc" as parameters, and return nil from the mocked call.
	mockDoer.On("Do", 1, "abc").Return(nil).Once()

	testUser.Use()

	mockDoer.AssertExpectations(t)
}