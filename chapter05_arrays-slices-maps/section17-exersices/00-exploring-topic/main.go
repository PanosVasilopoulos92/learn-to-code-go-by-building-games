package main

import (
	"fmt"
	"slices"
)

func main() {
	xi := []int{9, 2, 6, 4}
	fmt.Println(xi)
	xi = slices.Replace(xi, 1, 2, 10)
	fmt.Println(xi)
	index, ok := slices.BinarySearch(xi, 0)
	if !ok {
		fmt.Println("Value not found")
	}
	fmt.Println(xi[index])

	xi = slices.Clip(xi)
	fmt.Println(cap(xi))

	xi = slices.Delete(xi, 1, 2)
	fmt.Println(xi)

	slices.Sort(xi)
	fmt.Println(xi)

	minNum := slices.Min(xi)
	fmt.Println(minNum)

	maxNum := slices.Max(xi)
	fmt.Println(maxNum)

	// xj := xi // shallow copy
	// fmt.Println(xi, "\n", xj)
	// xj[0] = 12
	// fmt.Println(xi, "\n", xj)

	// deep copy
	xj := make([]int, len(xi))
	// for i, v := range xi {
	// 	xj[i] = v
	// }
	copy(xj, xi)
	fmt.Println(xj)

	xj[0] = 99
	fmt.Println(xi, "\n", xj)
}
