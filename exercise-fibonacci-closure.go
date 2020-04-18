package main

import "fmt"


// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib, fib1 := 0,0
	return func() int {
		if(fib == 0){
			fib = 1
			return fib1
		} else{
			prev := fib
			fib = fib + fib1
			fib1 = prev
			return fib
		} 
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
