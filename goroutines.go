package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg     = sync.WaitGroup{}
	dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6"}
)

// Add variable to store the results.
var results = []string{}

// Actually waitgroup is nothing but telling goroutine how many threads are there.
// that is go keyword before the func call.
// so we set the wg.add(//add how many routines you want to add.)
// once the routine are added just set the parent func to wait.
// for the routines to complete.
// Since this is the actual asychronous language not like the js

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	// Execute without waiting for the goroutine to complete.
	fmt.Printf("\n Total execution time: %v \n", time.Since(t0))

	wg.Wait()

	// The below line will wait till the goroutine gets completed.
	fmt.Printf("\n Total execution time: %v \n", time.Since(t0))
  fmt.Printf(" Results: %v \n", results)
}

func dbCall(i int) {
	// There may be problems when writing the data in a variable with multiple funcs.
	// so create read and write locks to avoid that.
	time.Sleep(time.Duration(2000) * time.Millisecond)
	// for that issue for writing only with one func at a time can be achieved by.
	// creating a read or write in another func which handles that for us.
	save(dbData[i], i)
	// Llog()
	// fmt.Printf("hello world %s \n", dbData[i])
	wg.Done()
}

// func Llog() {
// 	panic("unimplemented")
// }

func save(result string, i int) {
	results = append(results, dbData[i])
	fmt.Printf("hello world %s \n", dbData[i])
}
