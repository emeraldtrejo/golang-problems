package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Define an interface type called Animal which describes the methods of an animal.
//Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
type Animal interface {
	Speak()
	Move()
	Eat()
	GetName() string
}

type Bird struct {
	name string
}

type Cow struct {
	name string
}

type Snake struct {
	name string
}

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	sliceOfAnimals := []Animal{}

	fmt.Println("To begin, you can do either of the following types of queries:")
	fmt.Println("To create an animal for the database, you will need to use the following Syntax:")
	fmt.Println("newanimal (name of animal) (type of animal either cow, bird or snake)")
	fmt.Println("Upon creation of an animal, you can then query the animal(s) by doing the following:")
	fmt.Println("query (name of animal) (action of animal)")

	for {
		fmt.Print(">")
		userQuery, _ := inputReader.ReadString('\n')

		userQuery = userQuery[:len(userQuery)-2]
		sliceOfQuery := strings.Split(userQuery, " ")
		if len(sliceOfQuery) > 3 || len(sliceOfQuery) < 3 {
			fmt.Println("Invalid query, has to have 3 parts")
		}
		//check if first part is either newanimal or query of animal
		queryCommand := sliceOfQuery[0]
		switch queryCommand {
		case "newanimal":

			if sliceOfQuery[2] == "cow" {
				sliceOfAnimals = append(sliceOfAnimals, Cow{name: sliceOfQuery[1]})
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "snake" {
				sliceOfAnimals = append(sliceOfAnimals, Snake{name: sliceOfQuery[1]})
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "bird" {
				sliceOfAnimals = append(sliceOfAnimals, Bird{name: sliceOfQuery[1]})
				fmt.Println("Created it!")
			} else {
				fmt.Println("Invalid Query - not one of 3 options")
			}
		case "query":

			for _, Animal := range sliceOfAnimals {
				if Animal.GetName() == sliceOfQuery[1] {
					if sliceOfQuery[2] == "move" {
						Animal.Move()
					} else if sliceOfQuery[2] == "eat" {
						Animal.Eat()
					} else if sliceOfQuery[2] == "speak" {
						Animal.Speak()
					}
				}
			}

		default:
			fmt.Println("Invalid user query 3")

		}
	}

}

func (c Cow) GetName() string {
	return c.name
}

func (s Snake) GetName() string {
	return s.name
}

func (b Bird) GetName() string {
	return b.name
}

func (c Cow) Eat() {
	fmt.Printf("%s eats grass\n", c.name)
}

func (c Cow) Move() {
	fmt.Printf("%s walks\n", c.name)
}

func (c Cow) Speak() {
	fmt.Printf("%s moos\n", c.name)
}

func (s Snake) Eat() {
	fmt.Printf("%s eats mice\n", s.name)
}

func (s Snake) Move() {
	fmt.Printf("%s slithers\n", s.name)
}

func (s Snake) Speak() {
	fmt.Printf("%s hisses\n", s.name)
}

func (b Bird) Eat() {
	fmt.Printf("%s eats worms\n", b.name)
}

func (b Bird) Move() {
	fmt.Printf("%s flies\n", b.name)
}

func (b Bird) Speak() {
	fmt.Printf("%s peeps\n", b.name)
}
