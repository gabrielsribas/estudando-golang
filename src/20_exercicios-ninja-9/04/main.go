package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup
var mux sync.Mutex
var count int
const qtdgoroutines = 100

func main() {

	goroutines(qtdgoroutines)
	wg.Wait()
	fmt.Println("contador.: ", count, "\nQtd de Goroutines.: ", qtdgoroutines)
}

func goroutines(num int) {
	wg.Add(num)
	for i := 1; i <= num; i++ {
	   go func () {
	      mux.Lock()
	      v := count
	      runtime.Gosched()
	      v++
	      count = v
	      mux.Unlock()
	      wg.Done()
	   }()
	}
}
