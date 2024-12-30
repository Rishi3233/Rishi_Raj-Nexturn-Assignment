package employee

import (
	"errors"
	"fmt"
)

// Employee struct to hold employee details
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Slice to store employees
var employees []Employee

// AddEmployee function
func AddEmployee(id int, name string, age int, department string) error {
	// Check if ID is unique
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	// Validate age
	if age <= 18 {
		return errors.New("age must be greater than 18")
	}

	// Add employee
	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	return nil
}

// SearchEmployee function
func SearchEmployee(searchKey string) (*Employee, error) {
	for _, emp := range employees {
		if emp.Name == searchKey || fmt.Sprintf("%d", emp.ID) == searchKey {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

// ListEmployeesByDepartment function
func ListEmployeesByDepartment(department string) []Employee {
	var deptEmployees []Employee
	for _, emp := range employees {
		if emp.Department == department {
			deptEmployees = append(deptEmployees, emp)
		}
	}
	return deptEmployees
}

// CountEmployees function
func CountEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if emp.Department == department {
			count++
		}
	}
	return count
}
