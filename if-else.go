package main 
import (
	"fmt"
)

func main () {

     state_population := map[string] int{
		"up": 123,
		"uk": 111,
		"hp" : 100,
	 }
      
	 if pop, ok := state_population["up"]; ok {
		fmt.Println(pop)
	 }
	// fmt.Println(state_population)

}