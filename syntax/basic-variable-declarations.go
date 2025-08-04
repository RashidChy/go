package syntax

import "fmt"

func main() {
	// 1. Basic variable declarations
	var a int = 10         // integer
	var b float64 = 3.14   // floating-point number
	var c string = "hello" // string
	var d bool = true      // boolean

	// Short-hand declaration (only inside functions)
	x := 42
	y := "GoLang"

	// 2. Arrays (fixed size)
	var arr [3]int = [3]int{1, 2, 3}

	// 3. Slice (dynamic size, built on arrays)
	slice := []string{"apple", "banana"}
	slice = append(slice, "cherry")

	// 4. Map (key-value store, like dictionary)
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30

	// 5. Struct (custom data type)
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Charlie", Age: 40}

	// 6. Pointer (stores memory address)
	var p *int = &a // p points to a

	// Printing all values
	fmt.Println("Basic types:", a, b, c, d)
	fmt.Println("Shorthand vars:", x, y)
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Map:", m)
	fmt.Println("Struct:", person)
	fmt.Println("Pointer p points to:", *p) // *p means value at p
}
