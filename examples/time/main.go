package main

import (
	"fmt"
	"time"

	"github.com/aimo-x/algorithm"
)

func main() {
	var ds []time.Time
	nd := 30
	for i := 0; i < 30; i++ {
		ds = append(ds, time.Now().Add(time.Hour*time.Duration((24*i))))
	}
	// ds[4] = time.Now()
	var t algorithm.Time
	fmt.Println(t.ConsecutiveDays(ds, nd))
}
