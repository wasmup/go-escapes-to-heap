package main

import "fmt"

func main() {
	d := 2
	fun(d)
	fmt.Println("done")

}
func fun(e int) {
	a := 2
	fmt.Println(a)
	{
		b := e
		fmt.Println(b)
	}
	c := new(int)
	*c = 2
	fmt.Println(c)
	*c++
	fmt.Println(*c)
}
