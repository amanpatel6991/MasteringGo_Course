/*package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

        //fmt.Println(a.String1())

	fmt.Println(a, z)
}*/

package main

import "fmt"

type Car struct {
	Id    int
	Model string
}

func(c Car) String() string{
	return fmt.Sprintf("%v (%v-AD)" , c.Id , c.Model)
}

func main() {

	c:= Car{1 , "MG-002"}

	fmt.Println(c)

}
