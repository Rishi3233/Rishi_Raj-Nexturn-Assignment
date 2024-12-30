package main

import (
	"assignment_ems/employee"
	"github.com/gin-gonic/gin"
	"net/http"
	
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/employees", addEmployee)                  // Add an employee
	r.GET("/employees/:searchKey", searchEmployee)     // Search for an employee by ID or name
	r.GET("/employees/department/:department", listEmployeesByDepartment) // List employees by department
	r.GET("/employees/count/:department", countEmployees) // Count employees in a department

	// Start the server on port 8080
	r.Run(":8080")
}

// Handler to add an employee
func addEmployee(c *gin.Context) {
	var newEmployee struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Department string `json:"department"`
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Add employee
	err := employee.AddEmployee(newEmployee.ID, newEmployee.Name, newEmployee.Age, newEmployee.Department)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee added successfully"})
}

// Handler to search for an employee
func searchEmployee(c *gin.Context) {
	searchKey := c.Param("searchKey")

	emp, err := employee.SearchEmployee(searchKey)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emp)
}

// Handler to list employees by department
func listEmployeesByDepartment(c *gin.Context) {
	department := c.Param("department")

	employees := employee.ListEmployeesByDepartment(department)
	c.JSON(http.StatusOK, employees)
}

// Handler to count employees in a department
func countEmployees(c *gin.Context) {
	department := c.Param("department")

	count := employee.CountEmployees(department)
	c.JSON(http.StatusOK, gin.H{"department": department, "count": count})
}
