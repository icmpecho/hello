package hello

import "fmt"

// Greet return string "Hello" concatinated with name param
func Greet(name string) string {
	return fmt.Sprintf("Hello %s\n", name)
}
