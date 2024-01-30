package main

import (
	"fmt"
	"time"
  "sync"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id1"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
    wg.Add(1)
		go dbCall(i)
	}
  wg.Wait()
	fmt.Printf("\n Total execution time: %v \n", time.Since(t0))
}

func dbCall(i int) {
  time.Sleep(time.Duration(2000)*time.Millisecond)
  
  fmt.Printf("hello world %s \n", dbData[i])
  wg.Done()
}
