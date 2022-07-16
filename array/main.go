package main

import (
	"fmt"
	"math/rand"
	"time"

) 

func main() {
	var array [120]int
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < len(array); i++ {
		array[i] = r1.Intn(100)

	
	}


	fmt.Println(array)

}