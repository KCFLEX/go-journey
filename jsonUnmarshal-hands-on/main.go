package main

import (
	"encoding/json"
	"fmt"
)

// firstly create the structure
type Student struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	// now we are trying to unmarshal this json in a struct
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	var faculty []Student
	err := json.Unmarshal([]byte(s), &faculty)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(faculty)

	for i, v := range faculty {
		fmt.Printf("person no %v\n", i)
		fmt.Printf("name: %v, lastname: %v, age: %v", v.First, v.Last, v.Age)
		for _, saying := range v.Sayings {
			fmt.Println("\t", saying)

		}
	}

}
