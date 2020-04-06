//Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
//Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
// Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
//The first string is the name of an animal, either “cow”, “bird”, or “snake”.
//The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
//Your program should process each request by printing out the requested data.
package main

import (
	"fmt"
)

//You will need a data structure to hold the information about each animal.
//Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {

	var n string //name of animal
	var r string //reaction
	var a Animal //create instance of object

	for {
		fmt.Print(">")
		fmt.Scan(&n)
		fmt.Scan(&r)

		if n == "cow" {
			a = Animal{"grass", "walk", "moo"}
		} else if n == "bird" {
			a = Animal{"worms", "fly", "peep"}
		} else if n == "snake" {
			a = Animal{"mice", "slither", "hsss"}
		} else {
			fmt.Println("Incorrect input. Try again")
		}
		if r == "eat" {
			a.Eat()
		} else if r == "move" {
			a.Move()
		} else if r == "speak" {
			a.Speak()
		} else {
			fmt.Println("Incorrect input. Try again")
		}
	}

}

//Your program should call the appropriate method when the user makes a request.
// Make three methods called Eat(), Move(), and Speak().
//The receiver type of all of your methods should be your Animal type.
//The Eat() method should print the animal’s food,
func (name *Animal) Eat() {
	fmt.Println("I eat", name.food)
}

//the Move() method should print the animal’s locomotion,
func (name *Animal) Move() {
	fmt.Println("I", name.locomotion)
}

//and the Speak() method should print the animal’s spoken sound.
func (name *Animal) Speak() {
	fmt.Println("I", name.noise)
}
