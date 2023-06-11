// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Starting pointer demo for various cases!!!!");
	var name string = "Abhishek"
	var ptr_name *string = &name

	fmt.Println("Demostrating pointers and dereferencing");
	fmt.Println("Name is ", name)
	fmt.Println("ptr_name is ", ptr_name)
	fmt.Println("ptr_name derefenced=", *ptr_name)

	new_name := "Abhishek"
	new_ptr := &new_name
	fmt.Println("#######")
	fmt.Println("Shorthand initialization demo.")
	fmt.Println("new_name is ", new_name)
	fmt.Println("new_ptr is ", new_ptr)
	fmt.Println("new_name derefenced=", *new_ptr)
	fmt.Println("#######")
	
	arg1 := 10
	arg2 := 10
	fmt.Println("Pass by value and pass by reference demo");	
	fmt.Println("Init value of arg1", arg1)
	fmt.Println("Init value of arg2", arg2)
	demo_pass_by_value(arg1)
	fmt.Println("Value of arg1 after calling demo_pass_by_value", arg1)
	demo_pass_by_reference(&arg2)
	fmt.Println("Value of arg2 after calling demo_pass_by_reference", arg2)

}

func demo_pass_by_value(x int) {
	fmt.Println("Entering demo_pass_by_value with input", x)
	x = x + 10
	fmt.Println("New value of x in demo", x)
}

func demo_pass_by_reference(x *int) {
	fmt.Println("Entering demo_pass_by_value with input", x)
	fmt.Println("Entering demo_pass_by_value with input dereferenced value is", *x)
	*x++
	fmt.Println("New value of x in demo", *x)
}

