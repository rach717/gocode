package users_test

import (
	"simpleapp/mocks"
	"simpleapp/users"
	"testing"
)

func TestUserWithTestifyMock(t *testing.T) {
	mockDoer := &mocks.Doer{}

	testUser := &users.User{Doer: mockDoer}

	// Expect Do to be called once with 1 and "abc" as parameters, and return nil from the mocked call.
	mockDoer.On("CheckPalindrome", "abc").Return(false).Once()
	mockDoer.On("CheckPalindrome", "a").Return(true).Once()
	mockDoer.On("CheckPalindrome", "aaa").Return(true).Once()
	mockDoer.On("CheckPalindrome", "aaaa").Return(true).Once()
	mockDoer.On("CheckPalindrome", "abca").Return(false).Once()
	mockDoer.On("CheckPalindrome", "AbcBa").Return(true).Once()
	mockDoer.On("CheckPalindrome", "@ 2 @").Return(true).Once()
	mockDoer.On("CheckPalindrome", "@").Return(true).Once()
	mockDoer.On("CheckPalindrome", "@  #").Return(false).Once()
	mockDoer.On("CheckPalindrome", "").Return(false).Once()
	mockDoer.On("CheckPalindrome", "Abb a").Return(true).Once()
	mockDoer.On("CheckPalindrome", " ").Return(false).Once()
	testUser.Use()

	mockDoer.AssertExpectations(t)
}
