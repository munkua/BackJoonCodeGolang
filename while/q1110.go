package main

import "fmt"

func main() {

	var n, count, loopNum, remainder, divider, sum, temp int
	
	fmt.Scanf("%d\n", &n)
	temp = n
		
	if (n >= 0 && n <= 99) {

		if (n / 10 == 0) {

			n *= 10
			
		} else {

			for {

				if loopNum == n {
						
					break
						
				} else {

					divider = temp / 10
					remainder = temp % 10
					sum = divider + remainder
					loopNum = (remainder * 10) + (sum % 10)
					temp = loopNum
					fmt.Printf("%d\n", temp)
					count++
					
				}
					
			}
				
		}
		
	}
	
		
	
	
	fmt.Printf("%d\n", count)

}