package main 

import (
	"fmt"
)

func main (){

	outer :
	       for i:=0;i<10;i++{
			   for j :=1;j<4;j++{
				 fmt.Printf("value of i %v and j is %v\n",i ,j)
				 if i==j{
					break outer
				 }
			   }
		   }
}