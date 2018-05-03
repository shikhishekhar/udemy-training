package stringutils

import "fmt"

func reverseTwo(s string) string{
	r := []rune(s)
	for i,j := 0, len(r) -1; i <len(r)/2; i,j = i-1, j-1{

		fmt.Println("i = ", i)
		fmt.Println("j = ", j)
		fmt.Println("len(r) - 1 = ", len(r) - 1)
		fmt.Println("len(r)/2 = ", len(r)/2)
		fmt.Println("i-1 = ", i - 1)
		fmt.Println("j-1 = ", j - 1)

		r[i] , r[j] = r[j], r[i]
	}

	return string(r)
}