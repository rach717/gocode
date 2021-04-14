package users

import (
	"simpleapp/doer"
)

type User struct {
	Doer doer.Doer
}

func (u *User) Use() {
	u.Doer.CheckPalindrome("abc")
	u.Doer.CheckPalindrome("a")
	u.Doer.CheckPalindrome("aaa")
	u.Doer.CheckPalindrome("aaaa")
	u.Doer.CheckPalindrome("abca")
	u.Doer.CheckPalindrome("AbcBa")
	u.Doer.CheckPalindrome("@ 2 @")
	u.Doer.CheckPalindrome("@")
	u.Doer.CheckPalindrome("@  #")
	u.Doer.CheckPalindrome("")
	u.Doer.CheckPalindrome("Abb a")
	u.Doer.CheckPalindrome(" ")
}
