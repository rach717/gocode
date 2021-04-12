package users

import (
	"github.com/golang/mock/gomock"
	"neone/mocks"
	"testing"
)

func TestUser_GoMock_MissingCall(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)

	mockDoer.EXPECT().Do(1, "abc").Return(nil)
	mockDoer.EXPECT().Do(2, "def").Return(nil)

}

func TestUser_Testify_MissingCall(t *testing.T) {
	mockDoer := &mocks.Doer{}

	mockDoer.On("Do", 1, "abc").Return(nil)
	mockDoer.On("Do", 2, "def").Return(nil)

	mockDoer.AssertExpectations(t)
}
