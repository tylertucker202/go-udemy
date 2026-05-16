package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase
	// CalculateArea, UserInfo, NewHTTPRequest
	//Structs, interfaces, enums

	//snake_case
	// user_id, first_name, http_request

	//UPPERCASE
	// CONFIG_FILE_NAME

	// mixedCase
	// javaScript, htmlDocument, isValid
	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
