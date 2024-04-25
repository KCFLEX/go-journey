package main

import (
	"fmt"
)

//type person struct {
//	firstName               string
//	lastName                string
//	favoriteicecreamflavors []string
//}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	Make   string
	model  string
	doors  string
	colors string
}

func main() {
	/*	p1 := person{
			firstName: "James",
			lastName:  "bond",
			favoriteicecreamflavors: []string{
				"Vanilla", "Chocolate", "Strawberry", "Mint Chocolate Chip", "Cookies and Cream", "Rocky Road", "Butter Pecan",
			},
		}

		p2 := person{
			firstName: "Tom",
			lastName:  "Hanks",
			favoriteicecreamflavors: []string{
				"Coffee", "SaltedCaramel", "Matcha Green Tea",
			},
		}

		fmt.Println(p2)
		fmt.Println("------------")
		for _, v := range p1.favoriteicecreamflavors {
			fmt.Printf("%v\n", v)
		}
		fmt.Println("------------")
		c2 := map[string]person{
			p1.lastName: p1,
			p2.lastName: p2,
		}

		for _, v := range c2 {
			fmt.Printf("%v\n", v)
			for _, b := range v.favoriteicecreamflavors {
				fmt.Printf("%v - %v - %v\n", v.firstName, v.lastName, b)
			}
		}
	*/

	ferari := vehicle{
		engine: engine{
			electric: true,
		},
		Make:   "2017",
		model:  "model OS",
		doors:  "wing doors",
		colors: "yellow and black",
	}

	bugatti := vehicle{
		engine: engine{
			electric: false,
		},
		Make:   "2020",
		model:  "model S",
		doors:  "wing doors",
		colors: "black and gray",
	}

	fmt.Println(ferari, bugatti)
	fmt.Println(ferari.Make, bugatti.Make)

	rt := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "fredrick",
		friends: map[string]int{
			"arab":  4,
			"black": 6,
			"asian": 2,
		},
		favDrinks: []string{
			"Cola", "pepesi", "fanta",
		},
	}

	fmt.Println("-------------------")
	fmt.Println(rt)

}
