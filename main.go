package main

import (
	"fmt"
  "github.com/google/uuid"
)


func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Println(`The value of point ${*p}`, *p, i)

	// So basically what happening here is we get pointer of p
	// and the store the pointer of p in p.
	//
	// Which is simply p = pointer address which show memory address.
	// and *p shows the value in that memory addr.
	// *p means dereferencing the pointer to get value out of it.
	// &p shows where the pointer itself is stored.
	// p shows where the pointer points to.
	*p = 10
	fmt.Println("the value of *p and mem addr of i &i", *p, &i)

	p = &i

	// Since we referenced the p with i
	// Now the pointer points to the location of i.
	// So everytime i changes we change the the value of p
	fmt.Println(p, *p, &p)
	fmt.Println("update i")
	i = 20

	fmt.Println("the value of *p and i", *p, i)

	// Another way of accessing the pointer is using slice
	// Which is a golang way of saying dynamic array.
	// BE CAREFUL WITH THE SLICES SINCE SLICECOPY IN GO JUST TAKE THE pointer
	// SO UPDATING THE SLICE WILL RESULT IN UPDATING THE ORIGINAL

	k := []int{1, 2, 3}

	// %T shows the type of the value that's put out of the variable.
	// so p shows the *int32
	// and *p shows the int32

	fmt.Printf("k is %T %T\n", k, p)
	fmt.Printf("k is %T %T\n", k, *p)
}
