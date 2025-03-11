package main
import "fmt"

// add function returns the sum of two numbers
func add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println("Go CI/CD: Addition Program")
    fmt.Println("Sum of 5 + 3 =", add(5, 3))
}
