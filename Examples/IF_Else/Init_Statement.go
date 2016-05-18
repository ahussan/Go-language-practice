package main

func main() {
	b:=true

	if food := "Apple" ; b {
		println(food)
	}

	if food := "Orange" ; !b{
		println(food)
	}
}
