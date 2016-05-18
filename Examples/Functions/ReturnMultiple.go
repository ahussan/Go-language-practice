package main

import "fmt"

func main() {

	fmt.Print(returnmultiple("firstname", "Lastname "))

}

func returnmultiple(a string, b string)(string,string)  {
	Fformat:= fmt.Sprint(a, " ",b)
	Lformat:=fmt.Sprint(b," ",a)
	return Fformat, Lformat

}
