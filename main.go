package main

import "fmt"

func main() {
	fundData := getFundDataFromDir()

	fmt.Printf("fund data: %+v\n", fundData)
}
