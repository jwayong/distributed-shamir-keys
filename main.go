package main

import (
	"fmt"
	sssa "github.com/jwayong/sssa-golang"
)

func main() {
	
	sharesssa,_ := sssa.Create(3, 5, "januarwayong")
	fmt.Println("Encrypting text: januarwayong into 5 parts, minimum 3 parts to recover")
	fmt.Println("----------------------------------------------------------------------")
	for _, s := range sharesssa {
		fmt.Println(s)
	}

	//secretstring,_ := sssa.Combine(sharesssa[1:4])
	fmt.Println("Decrypting text: using minimum 3 parts")
	fmt.Println("----------------------------------------------------------------------")
	secretstring,_ := sssa.Combine([]string{sharesssa[0], sharesssa[2], sharesssa[3]})

	fmt.Println("Decrypted text: ")
	fmt.Println(secretstring)


}