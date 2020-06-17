package main

import (
	"fmt"
	"time"

	"github.com/aimo-x/algorithm"
)

func main() {
	var ds []time.Time
	nd := 2
	for i := 0; i < 5; i++ {
		ds = append(ds, time.Now().Add(time.Hour*time.Duration((24*i))))
	}
	ds[3] = time.Now()
	var t algorithm.Time
	fmt.Println(t.ConsecutiveDays(ds, nd))
}
