package main

import (
	"fmt"
	"github.com/nikosEfthias/FindShiftedNumbers/lib"
	"math"
	"runtime"
	"strconv"
	"time"
)

var limiter chan bool
var linesNum int = 2

func main() {
	var counter int
	fmt.Printf("\x1B[2J")
	go func() {
		for {
			fmt.Printf("\x1B[0H (%d)=> currentNum(%d)", runtime.NumGoroutine(), counter)
			time.Sleep(time.Millisecond * 5e2)
		}
	}()
	limiter = make(chan bool, *lib.J)
	for {
		limiter <- true
		go searchRange(counter)
		counter += *lib.IncrementBy
	}
}
func getPower(num int) int {
	return int(math.Floor(math.Log10(float64(num))))
}
func searchRange(a int) {
	var constNum bool
	var num int
	if num = getPower(a); num == getPower(a+(*lib.IncrementBy-1)) {
		constNum = true
	}
	for i := a; i < a+*lib.IncrementBy; i++ {
		if i < 100 {
			continue
		}
		if !constNum {
			num = getPower(i)
		}
		if i%int(math.Pow(10, float64(num))) == i / *lib.BaseNum && float64(i / *lib.BaseNum) == float64(i)/float64(*lib.BaseNum) {
			fmt.Printf("\x1B["+strconv.Itoa(linesNum)+"H FOUND %d /%d=%d \n", i, *lib.BaseNum, i / *lib.BaseNum)
			linesNum++
		}
	}
	<-limiter
}
