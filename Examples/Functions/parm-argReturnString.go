package main

import "fmt"

func main() {
	store:= joinString("firstName", "LastName")
	println(store)

}

func joinString(a string, b string) string  {
	res:= fmt.Sprint(a," ",b)
	return res
}
