package utils

import "fmt"

// HandleError function for reusable error handling
func HandleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
