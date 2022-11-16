package main

/**
* Importing fmt package for the sake of printing
*/
import (
	"fmt"
)

/**
* Animal is the name we want but since we are to use it as an interface, we will change the name into Animaler.
* Note that we will declare the methods to be used later here in this interface
 */
type Animaler interface {
	Eat()
	Move()
	Speak()
	Error()
}

/**
* A struct holding a string variable: SuperAnimals
 */
type SuperAnimals struct {
	locomotion string
}

/**
* An embedded struct holding content from another struct and two other string variables named Animals
 */
type Animals struct {
	SuperAnimals
	food  string
	sound string
}

/**
* Now we are indirectly implementing the Animaler interface without any keywords. We are about to define each method declared in the Animaler interface.
* This method will access the variable food in Animal class
 */
func (x Animals) Eat() {
	fmt.Println(x.food)
}

/**
* This method will access the variable locomotion in Animal class
 */
func (x Animals) Move() {
	fmt.Println(x.locomotion)
}

/**
* This method will access the variable sound in Animal class
 */
func (x Animals) Speak() {
	fmt.Println(x.sound)
}

func (x Animals) Error() {
	fmt.Println("Invalid query entered!")
}

/**
* Experiencing a dictionary / map in GO
* For the animal name as a key,
* that particular object is a value
 */
func mainOOP() {
	m := map[string]Animaler{
		"cow":   Animals{SuperAnimals{"walk"}, "grass", "moo"},
		"Cow":   Animals{SuperAnimals{"walk"}, "grass", "moo"},
		"Bird":  Animals{SuperAnimals{"fly"}, "worms", "peep"},
		"bird":  Animals{SuperAnimals{"fly"}, "worms", "peep"},
		"Snake": Animals{SuperAnimals{"slither"}, "mice", "hsss"},
		"snake": Animals{SuperAnimals{"slither"}, "mice", "hsss"},
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Enter animal name & query (eat / move / speak): ")
		fmt.Print(">")
		var animal, op string
		fmt.Scan(&animal)
		fmt.Print(">")
		fmt.Scan(&op)
		if op == "eat" {
			m[animal].Eat()
		} else if op == "move" {
			m[animal].Move()
		} else if op == "speak" {
			m[animal].Speak()
		} else {
			m[animal].Error()
		}
	}
}
