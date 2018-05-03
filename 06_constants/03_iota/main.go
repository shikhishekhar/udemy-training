package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1* 10)
	MB = 1 << (iota * 10) // 1 << (2* 10)
	GB = 1 << (iota * 10) // 1 << (3* 10)
	TB = 1 << (iota * 10) // 1 << (4* 10)
)

func main() {
	fmt.Println("KB - ", KB)
	fmt.Println("MB - ", MB)
	fmt.Println("GB - ", GB)
	fmt.Println("TB - ", TB)
}



















