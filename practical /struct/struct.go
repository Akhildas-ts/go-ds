package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Place string
}

func main() {

	student1 := Person{Name: "sumisha", Age: 19, Place: "palakkad"}
	student2 := &Person{Name: "akhil", Age: 18, Place: "palakkad"}

	fmt.Println("studentPointer", *student2)
	fmt.Println("student", student2)

	// c := 5
	// var p *int
	// p = &c

	// p

	teststruct(student2)
	teststruct2(student1)

	fmt.Println("student1", student1, "student 2", student2)
	// student3 :=
	// student4 :=
	// student5 :=

}

func teststruct2(t Person) {
	t.Name = "supersumisha"
}
func teststruct(t *Person) {

	t.Name = "superakhil"
}
