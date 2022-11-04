package main

import "fmt"

func main() {
	fund := newFundFromFile("ethicalGlobalFund.json")

	fmt.Printf("fund: %+v", fund)
}
