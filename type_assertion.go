package main 

import (
	"fmt"
)

type assert (i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

func main () {
	var s interface{}=27
	assert(s)
} 