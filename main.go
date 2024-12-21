package main

import "github.com/recktt77/employee-exe3.git/employee"

func main() {
	c := employee.Company{}
	c.AddPartTimeEmployee("Alixan", 2000, 5.2)
	c.AddFullTimeEmployee("Koliya", 89999)

	c.GetDetails()
}
