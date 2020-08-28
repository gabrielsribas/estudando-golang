package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("goroutines inicial.: ", runtime.NumGoroutine())
		
	bolastelecena := 100
	wg.Add(bolastelecena)
	for i := 1; i <= bolastelecena; i++ {
		go func() {
			runtime.Gosched()
			fmt.Println("Bola N˚.:", i)
			wg.Done()
		}()
		go fmt.Println("Enquanto estão tentando sortear a bola, vamos fazer outras coisas...")
	}
	fmt.Println("goroutines finais.: ", runtime.NumGoroutine())
	wg.Wait()
	
}
