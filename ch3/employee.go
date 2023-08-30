package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	employeeOne := employee{
		firstName: "n",
		lastName:  "n",
		id:        123,
	}

	emploeeTwo := employee{
		"a",
		"b",
		456,
	}

	var employeeThree = employee{}
	employeeThree.firstName = "c"
	employeeThree.lastName = "d"
	employeeThree.id = 789
	fmt.Println(employeeOne)
	fmt.Println(emploeeTwo)
	fmt.Println(employeeThree)
}
