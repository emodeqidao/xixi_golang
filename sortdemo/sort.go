package main

import (
	"sort"
	"fmt"
)

func main()  {
	sortdemo()
}

func sortdemo()  {
	a := []int{3, 6, 9, 3, 2, 1, 7}
	sort.Ints(a)

	for _, v := range a  {
		fmt.Println(v)
	}
}