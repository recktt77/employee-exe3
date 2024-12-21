package employee

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

type Company struct {
	FullTimeEmployees map[string]*FullTimeEmployee
	PartTimeEmployees map[string]*PartTimeEmployee
	nextID            uint64
}

func (c *Company) Init() {
	if c.PartTimeEmployees == nil {
		c.PartTimeEmployees = make(map[string]*PartTimeEmployee)
	}
	if c.FullTimeEmployees == nil {
		c.FullTimeEmployees = make(map[string]*FullTimeEmployee)
	}
}

func (c *Company) AddPartTimeEmployee(name string, hourlyRate uint64, hoursWorked float32) {
	c.nextID++
	employee := &PartTimeEmployee{
		ID:          c.nextID,
		Name:        name,
		HourlyRate:  hourlyRate,
		HoursWorked: hoursWorked,
	}
	c.PartTimeEmployees[name] = employee
}

func (c *Company) AddFullTimeEmployee(name string, salary uint32) {
	c.nextID++
	employee := &FullTimeEmployee{
		ID:     c.nextID,
		Name:   name,
		Salary: salary,
	}
	c.FullTimeEmployees[name] = employee
}

func (c *Company) GetDetails() string {
	result := "Company Employees:\n"

	if len(c.FullTimeEmployees) > 0 {
		result += "FullTimeEmployees:\n"
		for _, employee := range c.FullTimeEmployees {
			result += fmt.Sprintf("id: %v, name: %v, salary: %v\n", employee.ID, employee.Name, employee.Salary)
		}
	} else {
		result += "No FullTimeEmployees\n"
	}

	if len(c.PartTimeEmployees) > 0 {
		result += "PartTimeEmployees:\n"
		for _, employee := range c.PartTimeEmployees {
			result += fmt.Sprintf("id: %v, name: %v, Hourly Rate: %v, Hours Worked: %v\n", employee.ID, employee.Name, employee.HourlyRate, employee.HoursWorked)
		}
	} else {
		result += "No PartTimeEmployees\n"
	}

	return result
}
