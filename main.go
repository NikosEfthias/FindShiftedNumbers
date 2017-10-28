package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const Num = 57
const incrementBy int = 5e2

var limiter chan bool

func main() {
	var counter int
	go func() {
		for {
			fmt.Printf("\r (%d)=> currentNum(%d)", runtime.NumGoroutine(), counter)
			time.Sleep(time.Millisecond * 5e2)
		}
	}()
	limiter = make(chan bool, 200)
	for {
		limiter <- true
		go searchRange(counter)
		counter += incrementBy
	}
}
func getPower(num int) int {
	return int(math.Floor(math.Log10(float64(num))))
}
func searchRange(a int) {
	var constNum bool
	var num int
	if num = getPower(a); num == getPower(a+(incrementBy-1)) {
		constNum = true
	}
	for i := a; i < a+incrementBy; i++ {
		if i < 100 {
			continue
		}
		if !constNum {
			num = getPower(i)
		}
		if i%int(math.Pow(10, float64(num))) == i/Num && float64(i/Num) == float64(i)/float64(Num) {
			fmt.Printf("\n\nFOUND %d /57=%d \n", i, i/Num)
		}
	}
	<-limiter
}
