package metrics

import (
	"fmt"
	"time"
)

func PrintResultAndTime[A any](name string, f func() (A, error)) {
	start := time.Now()
	result, err := f()
	total := time.Since(start)
	if err != nil {
		fmt.Printf("%s(%dms)=ERROR %v\n", name, total.Milliseconds(), err)
	} else {
		fmt.Printf("%s(%dms)=%v\n", name, total.Milliseconds(), result)
	}
}
