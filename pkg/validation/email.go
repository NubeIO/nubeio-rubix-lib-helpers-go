package validation

import (
	"errors"
	"github.com/mcnijman/go-emailaddress"
)

/*
	_email, err := emailaddress.Parse("1a-foobar.com")
	if err != nil {
		fmt.Println("invalid email")
	} else {
		fmt.Println(_email.LocalPart) // foo
		fmt.Println(_email.Domain)    // bar.com
		fmt.Println(_email)           // foo@bar.com
		fmt.Println(_email.String())  // foo@bar.com
	}
*/

func IsEmail(address string) (bool, error) {
	_, err := emailaddress.Parse(address)
	if err != nil {
		return false, errors.New("invalid subnet address, 255.255.255.0")
	}
	return true, nil
}
