package main

import "fmt"

func main() {

	var a, b, c, count int32
	fmt.Scanf("%d %d %d\n", &a, &b, &c)
	
	
	if b > c || b == c {

		count--
	
	} else {

			count = a / (c - b) + 1
		
		}
	
	fmt.Printf("%d\n", count)
	
	}
	

