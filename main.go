package main

import (
	"fmt"
	sssa "github.com/SSSaaS/sssa-golang"
)

func main() {
	
	sharesssa,_ := sssa.Create(3, 5, "januarwayong")

	for _, s := range sharesssa {
		fmt.Println(s)
	}
	fmt.Println("")
	for _, s := range sharesssa[1:4] {
		fmt.Println(s)
	}

	//secretstring,_ := sssa.Combine(sharesssa[1:4])
	secretstring,_ := sssa.Combine([]string{sharesssa[0], sharesssa[2], sharesssa[3]})


	fmt.Println(secretstring)


}