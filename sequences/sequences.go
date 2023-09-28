package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(i)
	}
	fmt.Println(slice)
}

func mapArray(f func(a int) int, array [3]int) {
	for i := 0; i < len(array); i++ {
		array[i] = f(i)
	}
	fmt.Println(array)
}

func main() {

}
