package main

import "fmt"

func mars_age(a int) int {
	days := a * 365
	mars_years := days / 687
	return mars_years
}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}

//sample in 42 sample out 22
//earth year 365
//mars year 687
