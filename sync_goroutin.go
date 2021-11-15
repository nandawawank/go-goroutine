package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// fungsi add untuk memberitahu
	// ada 2 goroutine yang akan dijalankan secara syncronus
	// wait berfungsi untuk menunggu 2 goroutine yang sedang berjalan
	wg.Add(2)
	go SayHello("Nanda")
	go SayHello("Wawan")
	wg.Wait()

	wg.Add(1)
	go SayHello("Kurniawan")
	wg.Wait()
}

func SayHello(name string) {
	// defer berfungsi untuk menahan sintak agar dijalankan terakhir
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Hello " + name)
		time.Sleep(time.Millisecond * 100)
	}
}
