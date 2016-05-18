package main

func main() {

	i:=1

	for i<=50 {
		if i % 2 == 0 {
			println(i, " is a even Number")
		}else if i % 2 != 0 {
			println(i, " is an odd number")
		}else {
			println(i)
		}
		i++
	}



}
