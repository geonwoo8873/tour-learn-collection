package main

// import "fmt"

// func <funcation_name> (<parameter_name> <parameter_type>) <return data type> {...}
func process(s string) string {
	// Print the provided string to the console
	return "String : " + s
}

func main() {

	/*
	// Error :
	// cannot use 432 (untyped int constant) as string value in argument to process
	// The not created binary representation of the integer cannot be used as a string
	*/
	// result := process(432)

	result := process("Hello")
	println(result)
}