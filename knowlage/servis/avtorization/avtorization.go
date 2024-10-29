package avtorization

import (
	"fmt"
	"net/http"
)

type Avtorization struct {
}

func NewAvtorization() Avtorization {
	return Avtorization{}
}

func (das Avtorization) Registration(password string, login string, phoneNumber int) {
	Registration(password, login, phoneNumber)
}

// ---
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
func Registration(password string, login string, phoneNumber int) {
	//var polzovatel Guest= Guest{}
	polzovatel := Guest{
		password:    password,
		name:        login,
		phoneNumber: phoneNumber,
	}
	MnogoGuestov = append(MnogoGuestov, polzovatel)

}
func VxodKeys(w http.ResponseWriter, r *http.Request) {
	v := func(string, string) {
		Registration(password string, login string)
	}
	v()

}
