package main

import (
	"fmt"
	"log"
	"strconv"
)

/*func (r reciever) identifier(p parameter(s)) (return(s)) {
...
}*/

type dog struct {
	name string
}

// embedded struct
type policeDog struct {
	dog
	ltk bool
}

// this is an interface: an interface allows values to have m ore than one type
type Wolf interface {
	sound()
}

func dogBark(w Wolf) {
	w.sound()
}

func main() {
	/*defer says don't run this function untill the outer surrounding function is done excuting */
	defer shout()
	answer()

	/////////////////////////////////////////////////////////////////////
	fmt.Println("---------")
	germanSheperd := policeDog{
		dog: dog{
			name: "scott",
		},
		ltk: true,
	}
	chiwawa := dog{
		name: "chi chi",
	}
	rottweiler := dog{
		name: "django",
	}
	// they both inherited the sound method
	//germanSheperd.sound()
	//chiwawa.sound()
	//rottweiler.sound()

	dogBark(rottweiler)
	dogBark(germanSheperd)
	dogBark(chiwawa)

	fmt.Println("----------------------")
	//stringer call
	b := book{
		title: "james bond",
	}

	var n count = 42

	logInfo(b)
	logInfo(n)

}

func shout() {
	fmt.Println("are you mad?!")
}

func answer() {
	fmt.Println("no")
}

// methods in go
// what you see below is a method
func (d dog) /*(d dog) is a reciever*/ sound() {
	fmt.Println(d.name, "barked woof woof")
}
func (pd policeDog) /*(d dog) is a reciever*/ sound() {
	fmt.Println(pd.name, "barked grrrrr woof woof woof")
}

// stringer

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

// wrapper function
func logInfo(s fmt.Stringer) {
	/*log */ log.Println("LOG FROM 138", s.String())
}
