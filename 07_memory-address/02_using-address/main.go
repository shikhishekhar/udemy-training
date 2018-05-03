package main

import "fmt"

func main() {
	var meters float64

	fmt.Print("Enter meters - ")
	fmt.Scan(&meters)
	yards := meters * 10
	fmt.Println("Meters - ", meters)
	fmt.Println("Yards - ", yards)
}
