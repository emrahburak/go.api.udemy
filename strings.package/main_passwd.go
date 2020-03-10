package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Rastgele string oluşturma

type Option struct {
	length    int
	upperCase bool
	loverCase bool
	numbers   bool
	special   bool
}

var source = rand.NewSource(time.Now().UnixNano())

// const _charset = "abcdefghijklmnoprstuwxyzABCDEFGHIJKLMNOPRSTUWXYZ0123456789"

const _charsetLovercase = "abcdefghijklmnoprstuwxyz"
const _charsetUppercase = "ABCDEFGHIJKLMNOPRSTUWXYZ"
const _charsetNumbers = "0123456789"
const _charsetSpecial = "[]^$'%)("

func generatePassword(opt Option) (string, error) {
	x := make([]byte, opt.length)
	charset := "."

	if opt.upperCase {
		charset += _charsetUppercase
	}

	if opt.loverCase {
		charset += _charsetLovercase
	}

	if opt.numbers {
		charset += _charsetNumbers
	}

	if opt.special {
		charset += _charsetSpecial
	}

	if charset == "." {
		return "NON-GENERETED", errors.New("Herhangi bir karakter seçimi yapmak zorundasınız")
	}

	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]

	}
	return string(x), nil
}

func main() {

	password, err := generatePassword(Option{length: 15, upperCase: true, loverCase: true, special: true, numbers: true})

	fmt.Println(password)
	if err != nil {
		fmt.Println(err.Error)
	}

}
