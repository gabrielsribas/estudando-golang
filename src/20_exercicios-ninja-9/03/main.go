package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup
var count int

func main() {

	gos := 10000
	goroutines(gos)
	runtime.Gosched()
	fmt.Println("contador.: ", count, "\nQtd de Goroutines.: ", runtime.NumGoroutine())
	wg.Wait()
}

func goroutines(num int) {
	j := num
	wg.Add(j)
	for i := 1; i <= j; i++ {
	   go func (x int) {
	      v := count
	      runtime.Gosched()
	      v++
	      count = v
	      wg.Done()
	   }(i)
	}
}
