package main

import (
	"fmt"
)

func main(){
   fmt.Println("prime")

   for x := 0; x <= 100; x++ {
	   if isPrime(x) {
		fmt.Print(x)
		fmt.Print(" ")
	   }
   }
}

func isPrime(n int) bool {
	
	if n == 1 {
		return false
	} else if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

    for x := 3; x*x <=n; x+=2 {
		if n%x == 0 {
			return false;
		}		
	}

   return true
  }
