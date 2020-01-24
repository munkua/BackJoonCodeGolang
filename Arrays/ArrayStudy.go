package main

import (
	"fmt"
	"crypto/sha256"
)

func visible(*ptr [32]byte) {

	for i := range ptr {

		fmt.Printf("%d ", i)
		
	}
	
	fmt.Println()

}

func main() {

	digest1 := sha256.Sum256([]byte("has"))
	digest2 := sha256.Sum256([]byte("had"))
	
	visible(digest1)
	visible(digest2)

}

