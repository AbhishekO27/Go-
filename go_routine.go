package main 

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("hello world")
}

func main (){
	go hello();
    time.Sleep(1*time.Second)
	//time.Sleep(1 * time.Second)
	fmt.Printf("welcome all")
}