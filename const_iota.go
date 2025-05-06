package main 
import "fmt"

const (
	
	error = iota
	catspecialist
	dogspecialist
	snakespecialis
)

const (
	a2=iota
)

func main (){

	var specialist int 
	fmt.Printf("%v", specialist==catspecialist)
}