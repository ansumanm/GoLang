package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r *Rectangle) Scale(factor int) {
	r.width = r.width * factor
	r.height = r.height * factor
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 5, height: 3}

	// Method with a value receiver
	area := rect.Area()
	fmt.Println("Area:", area) // Output: 30 (Unchanged)

	// Method with a pointer receiver
	rect.Scale(2)
	fmt.Println("Width:", rect.width)   // Output: 10 (Changed)
	fmt.Println("Height:", rect.height) // Output: 6 (Changed)

	// Method with a value receiver
	area = rect.Area()
	fmt.Println("Area:", area) // Output: 60 (Changed)
}
