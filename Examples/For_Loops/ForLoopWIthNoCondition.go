package main

import "fmt"

func main() {
	i:=0
	// this for loop has no condition
	for  {
		fmt.Println(i)
		if i==20 {
			break
		}
		i++
	}
}
