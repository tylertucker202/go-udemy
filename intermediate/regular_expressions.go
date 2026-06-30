package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I am Great\"")

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	//Match
	email1 := "user@email.com"
	email2 := "user.email.com"

	fmt.Printf("%s is valid: %t\n", email1, re.MatchString(email1))
	fmt.Printf("%s is valid: %t\n", email2, re.MatchString(email2))

	result := re.FindAllString("My email is: tylertucker202@gmail.com", -1)
	fmt.Println(result)

	//Capturing groups
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2024-06-01"
	matches := re.FindStringSubmatch(date)
	if len(matches) > 0 {
		fmt.Printf("Year: %s, Month: %s, Day: %s\n", matches[1], matches[2], matches[3])
	} else {
		fmt.Println("No match found")
	}

	//Target string
	str := "Hello world"
	re = regexp.MustCompile(`[aeiou]`)
	result2 := re.ReplaceAllString(str, "*")
	fmt.Println(result2)

	// i flag for case-insensitive matching
	// m flag for multi-line matching
	// s flag for dotall mode (dot matches newline)
	// U flag for ungreedy matching

	re = regexp.MustCompile(`(?i)go`)
	fmt.Println(re.MatchString("Go is great"))

}
