package main
 
import "fmt"


func main() {
	
	var n, x, elem int
	_,err := fmt.Scanf("%d %d\n", &n ,&x)
	
	if err == nil {
		
		for i := 0; i < n; i ++ {
		
			fmt.Scanf("%d", &elem)
		
			if elem < x {

				fmt.Printf("%d ", elem)
		
			}
			
		}
		
	}
	
}