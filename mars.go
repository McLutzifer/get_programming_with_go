// My weight loss program.
package main

import "fmt"

//main is the functionwhere it all begins.
func main() {

	//Print print //Println line //Printf insert values
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(93 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(30 * 365 / 687)
	fmt.Print(" years old.")

	fmt.Println("My weight n Mars is", 149.0*0.3783)

	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*465/687)

	fmt.Printf("My weigh on the surface of %v is %v lbs.\n\n", "Earth", 149.0)

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
