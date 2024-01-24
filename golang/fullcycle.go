package main

import (
	"fmt"
	"time"
)

func main() {
    for i:=1; i>0; i++ {
		fmt.Println("Full Cycle Rocks!!")
		time.Sleep(10 * time.Second)
	}
}