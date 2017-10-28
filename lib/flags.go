package lib

import (
	"flag"
)

//FLAGS
var (
	//J is the parallelization factor
	J *int

	//BaseNum is the number that is as base num
	BaseNum *int
	//IncrementBy is the number each batch is incremented by
	IncrementBy *int
)

func init() {
	J = flag.Int("j", 200, "Parallelization factor")
	BaseNum = flag.Int("base", 7, "Number to check powers of")
	IncrementBy = flag.Int("i", 5000, "How much to increment each batch")
	//<<		-------------		>>
	if !flag.Parsed() {
		flag.Parse()
	}
}
