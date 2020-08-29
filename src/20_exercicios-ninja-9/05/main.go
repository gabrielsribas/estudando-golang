package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"
)

var wg sync.WaitGroup
var mux sync.Mutex
var count int32
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
	      v := &count
	      atomic.AddInt32(v, 1)
	      fmt.Println(*v)
	      runtime.Gosched()
	      wg.Done()
	   }()
	}
}
