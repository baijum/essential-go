package main

import "fmt"

type UnauthorizedError struct {
	UserID string
}

func (e UnauthorizedError) Error() string {
	return "User not authorised: " + e.UserID
}

func SomeAction() error {
	return UnauthorizedError{"jack"}
}

func main() {
	err := SomeAction()
	if err != nil {
		fmt.Println(err)
	}
}
