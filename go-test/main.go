package main

import (
	"fmt"

	"github.com/TeerapatChan/go-example/teerapat"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello World!")
	teerapat.SayHello()
	teerapat.SayTest()
	// teerapat.SayLopster()
	fmt.Println("UUID: ", id)
}

func myArray() {
	var myArray [3]int // An array of 3 integers
	myArray[0] = 10    // Assign values
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println(myArray) // Output: [10 20 30]
}

func mySlice() {
	mySlice := []int{10, 20, 30, 40, 50} // A slice of integers

	fmt.Println(mySlice)          // Output: [10 20 30 40 50]
	fmt.Println(len(mySlice))     // Length of the slice: 5
	fmt.Println(cap(mySlice))     // Capacity of the slice: 5

	// Slicing a slice
	subSlice := mySlice[1:3]      // Slice from index 1 to 2
	fmt.Println(subSlice)         // Output: [20 30]
}

func myMap() {
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	// Access and print a value for a key
	fmt.Println("Apples:", myMap["apple"])

	// Update the value for a key
	myMap["banana"] = 12

	// Delete a key-value pair
	delete(myMap, "orange")

	// Iterate over the map
	for key, value := range myMap {
	fmt.Printf("%s -> %d\n", key, value)
	}

	// Checking if a key exists
	val, ok := myMap["pear"]
	if ok {
	fmt.Println("Pear's value:", val)
	} else {
	fmt.Println("Pear not found inmap")
	}
}

type Student struct {
	Firstname string
	Lastname  string
	Weight  int
	Height  int
	Grade   string
}

func myStruct() {
	// Create an instance of the Student struct
	var student1 Student
	student1.Firstname = "John"
	student1.Lastname = "Doe"
	student1.Weight = 60
	student1.Height = 180
	student1.Grade = "F"

	// Print struct values
	fmt.Println(student1)
	fmt.Println("Fullname:", student1.FullName())
}

// Method with a receiver of type Student
// This method returns the full name of the student
func (s Student) FullName() string {
	return s.Firstname + " " + s.Lastname
}


// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Person struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p Person) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func myInterface() {
	dog := Dog{Name: "Buddy"}
	person := Person{Name: "Alice"}

	makeSound(dog)
	makeSound(person)
}