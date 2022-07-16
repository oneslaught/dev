package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var array [10]int
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < len(array); i++ {
		array[i] = r1.Intn(10)
	}

	fmt.Println(array)

	mapa := make(map[int]int)
	for i := 0; i < len(array); i++ {
		mapa[array[i]] += 1
	}
	fmt.Println(mapa)

	var maximum int = 0
	var mostFrequendNumbers []int
	for key, value := range mapa {
		if maximum < value {
			maximum = value
			mostFrequendNumbers = make([]int, 0)
			mostFrequendNumbers = append(mostFrequendNumbers, key)
		} else if maximum == value {
			mostFrequendNumbers = append(mostFrequendNumbers, key)
		}
	}
	fmt.Println("mostFrequendNumbers is ", mostFrequendNumbers, ", appeared ", maximum, " times")
}
