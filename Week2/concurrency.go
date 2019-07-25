package main

import (
	"fmt"
	"sync"
)

func spread(main chan int, a chan int, b chan int, c chan int, wg *sync.WaitGroup) {
	for {
		value := <-main
		a <- value
		b <- value
		c <- value
		fmt.Printf("Insertion to main: %d", value)
		wg.Done()
	}

}

func printChannel(c chan int, f func(int) int) {
	for {
		value := <-c
		fmt.Println(f(value))
	}
}

/*
func main() {
	var wg sync.WaitGroup
	main := make(chan int)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	go spread(main, a, b, c, &wg)
	defer wg.Wait()
	go printChannel(a, func(i int) int {
		return i * 2
	})
	go printChannel(a, func(i int) int {
		return i * 3
	})
	go printChannel(a, func(i int) int {
		return i * 4
	})
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		main <- i
	}
}
*/
