package main

import "fmt"

func main() {
	
	var a, b int

	for {

		fmt.Scanf("%d %d\n", &a, &b)
		
		if a == 0 && b == 0 {

		break
		
		} else {

			fmt.Printf("%d\n", a + b)
		
		}
	
	}

}