package main

import "fmt"

func double(x int) (result int) {

	defer func() {
		result += x
		fmt.Printf("double(%d)=%d \n", x, result)

	}()

	return x + x

}

func main() {

	fmt.Println(double(20))

}