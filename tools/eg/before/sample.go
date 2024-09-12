package sample

import (
	"errors"
	"fmt"
)

func sample() {
	msg := "something went wrong"
	err := errors.New("error: " + msg)
	fmt.Println(err)
}
