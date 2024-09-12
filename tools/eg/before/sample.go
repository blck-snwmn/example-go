package sample

import (
	"fmt"
)

func sample() {
	msg := "something went wrong"
	err := fmt.Errorf("%s", "error: "+msg)
	fmt.Println(err)
}
