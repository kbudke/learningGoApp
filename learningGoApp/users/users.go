package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func users() {
	userPW := `password123`
	bcryptPW, err := bcrypt.GenerateFromPassword([]byte(userPW), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please insert password")
	fmt.Println(userPW)
	fmt.Println(bcryptPW)

	loginPword1 := `password123`

	err = bcrypt.CompareHashAndPassword(bcryptPW, []byte(loginPword1))
	if err != nil {
		fmt.Println("You Can't Login")
		return
	}
	fmt.Println("Your logged in!")
}
