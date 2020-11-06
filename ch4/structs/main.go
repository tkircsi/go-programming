package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	t, _ := time.Parse("2006-01-02", "1970-11-22")
	dilbert := Employee{
		10,
		"Dilbert",
		"Budapest",
		t,
		"Worker",
		20000,
		0,
	}

	fmt.Println(dilbert)

	dilbert.Salary -= 5000

	position := &dilbert.Position
	*position = "Product Manager"

	var employee *Employee = &dilbert
	employee.Address = "Debrecen"
	fmt.Println(dilbert)

	fmt.Printf("%p\n", &dilbert)
	dilbert = promoteEmployee(&dilbert, "CEO")
	fmt.Println(dilbert)
	fmt.Printf("%p\n", &dilbert)

	fmt.Printf("%p\n", &dilbert)
	promoteEmployee2(&dilbert, "CFO")
	fmt.Println(dilbert)
	fmt.Printf("%p\n", &dilbert)

}

func promoteEmployee(e *Employee, postition string) Employee {
	// (*e).Position = position
	e.Position = postition
	return *e
}

func promoteEmployee2(e *Employee, postition string) {
	e.Position = postition
}
