package main

import "fmt"

func main() {
	var x [10]string

	fmt.Println("x: ", x)
	fmt.Println("len(x): ", len(x))
	fmt.Println("x[9]: ", x[9])

	for i := 0; i < 10; i++ {
		fmt.Println("i ===", i)
		x[i] = fmt.Sprint(i)
		fmt.Println("x[i] ====", x[i])
	}
	fmt.Println("x: ", x)
	fmt.Println("len(x): ", len(x))
	fmt.Println("x[9]: ", x[9])
}
