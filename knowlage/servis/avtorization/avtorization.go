package avtorization

import "fmt"

type Guest struct {
	name         string
	ID           int
	balanse      int
	password     string
	registeredAt int
	pole         string
	phoneNumber  int
	status       string
}

var MnogoGuestov []Guest = []Guest{}

func CreateUser(Name string, ps string, PNumber int) {
	MnogoGuestov = append(MnogoGuestov, Guest{name: Name, password: ps, phoneNumber: PNumber})
	fmt.Println(MnogoGuestov)
}
