package main


func main() {

	println(addition(100, 399))
	println(subtract(455,509))
}

func addition (x int, y int) int{
	z := x + y
	return z
}

func subtract(x int, y int) int  {
	result:= x-y
	return result
}
