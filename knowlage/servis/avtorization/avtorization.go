package avtorization

import (
	"fmt"
	//"net/http"
)

type Avtorization struct {
}

func NewAvtorization() Avtorization {
	return Avtorization{}
}

func (das Avtorization) Registration(login string, phoneNumber string) {
	Registration(login, phoneNumber)
}

// ---
type Guest struct {
	name         string
	ID           int
	balanse      int
	password     string
	registeredAt int
	pole         string
	phoneNumber  string
	status       string
}

var MnogoGuestov []Guest = []Guest{}

func CreateUser(Name string, PNumber string) {
	MnogoGuestov = append(MnogoGuestov, Guest{name: Name, phoneNumber: PNumber})
	fmt.Println(MnogoGuestov)
}
func Registration(login string, phoneNumber string) {
	//var polzovatel Guest= Guest{}
	polzovatel := Guest{
		name:        login,
		phoneNumber: phoneNumber,
	}
	MnogoGuestov = append(MnogoGuestov, polzovatel)

}

// func VxodKeys(w http.ResponseWriter, r *http.Request) {
// 	v := func(string, string) {
// 		Registration(password string, login string)
// 	}
// 	v()

// }
