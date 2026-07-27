package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	var pool = sync.Pool{
		// New: func() interface{} {
		// 	fmt.Println("Creating a new person.")
		// 	return &person{}
		// },
	}

	pool.Put(&person{name: "John", age: 81})

	//Get an object from the pool
	person1 := pool.Get().(*person)

	fmt.Println("got person:", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)
	pool.Put(person1)

	person2 := pool.Get().(*person)
	fmt.Println("got person2:", person2)

	//person3 := pool.Get().(*person)
	person3 := pool.Get()
	if person3 != nil {
		person3.(*person).name = "Jane"
		fmt.Println("got another person:", person3)
		pool.Put(person3)
	} else {
		fmt.Println("Sync.Pool is empty")

	}

	//returning to pool again
	pool.Put(person2)
	fmt.Println("Returned persons to pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got person:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got person:", person5)
}
