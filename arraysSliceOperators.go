// Go Arrays,Go Slices,Go Operators
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//var array_name = [length]datatype{values} // here length is defined
	//var array_name = [...]datatype{values} // here length is inferred
	//array_name := [...]datatype{values} // here length is inferred

	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars, "\n")

	prices := [...]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	prices_ := [3]int{10, 20, 30}

	prices_[2] = 50
	fmt.Println(prices_)

	arr11 := [5]int{}              //not initialized
	arr21 := [5]int{1, 2}          //partially initialized
	arr31 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr11)
	fmt.Println(arr21)
	fmt.Println(arr31)

	//Initialize Only Specific Elements
	arr12 := [5]int{1: 10, 2: 40}

	fmt.Println(arr12)

	arr13 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr23 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr13))
	fmt.Println(len(arr23))

	//Go Slices

	/*
		Slices are similar to arrays, but are more powerful and flexible.

		Like arrays, slices are also used to store multiple values of the same type in a single variable.

		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

		In Go, there are several ways to create a slice:

		    Using the []datatype{values} format
		    Create a slice from an array
		    Using the make() function
	*/
	//slice_name := []datatype{values}

	/*
			In Go, there are two functions that can be used to return the length and capacity of a slice:

		    len() function - returns the length of the slice (the number of elements in the slice)
		    cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)???????????????????

	*/

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	/*var myarray = [length]datatype{values} // An array
	myslice := myarray[start:end] // A slice made from the array
	*/

	arr15 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr15[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	/*
		Create a Slice With The make() Function
		The make() function can also be used to create a slice.
		slice_name := make([]type, length, capacity)
	*/

	myslice12 := make([]int, 5, 10)
	fmt.Printf("myslice12 = %v\n", myslice12)
	fmt.Printf("length = %d\n", len(myslice12))
	fmt.Printf("capacity = %d\n", cap(myslice12))

	myslice22 := make([]int, 5)
	fmt.Printf("myslice22 = %v\n", myslice22)
	fmt.Printf("length = %d\n", len(myslice22))
	fmt.Printf("capacity = %d\n", cap(myslice22))

	//Go Access, Change, Append and Copy Slices
	//Go Access, Change, Append and Copy Slices

	prices7 := []int{10, 20, 30}

	fmt.Println(prices7[0])
	fmt.Println(prices7[2])

	prices5 := []int{10, 20, 30}
	prices5[2] = 50
	fmt.Println(prices5[0])
	fmt.Println(prices5[2])
	fmt.Println(prices5)

	//slice_name = append(slice_name, element1, element2, ...)

	myslice17 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice17)
	fmt.Printf("length = %d\n", len(myslice17))
	fmt.Printf("capacity = %d\n", cap(myslice17))

	myslice17 = append(myslice17, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice17)
	fmt.Printf("length = %d\n", len(myslice17))
	fmt.Printf("capacity = %d\n", cap(myslice17))

	myslice19 := []int{1, 2, 3}
	myslice29 := []int{4, 5, 6}
	myslice39 := append(myslice19, myslice29...)
	// myslice39 := append(myslice19, myslice29) error disse

	fmt.Printf("myslice39=%v\n", myslice39)
	fmt.Printf("length=%d\n", len(myslice39))
	fmt.Printf("capacity=%d\n", cap(myslice39))

	arr44 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice31 := arr44[1:5]                // Slice array
	fmt.Printf("myslice31 = %v\n", myslice31)
	fmt.Printf("length = %d\n", len(myslice31))
	fmt.Printf("capacity = %d\n", cap(myslice31))

	myslice31 = arr44[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice31 = %v\n", myslice31)
	fmt.Printf("length = %d\n", len(myslice31))
	fmt.Printf("capacity = %d\n", cap(myslice31))

	myslice31 = append(myslice31, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice31 = %v\n", myslice31)
	fmt.Printf("length = %d\n", len(myslice31))
	fmt.Printf("capacity = %d\n", cap(myslice31))

	/*
			Memory Efficiency

		 When using slices, Go loads all the underlying elements into the memory.

		If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

		The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
	*/

	/*copy(dest, src)
	The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	// copy(numbersCopy, neededNumbers)
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", neededNumbers)
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

//https://www.w3schools.com/go/go_operators.php
