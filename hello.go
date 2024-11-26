package hello

import "fmt"

func GetHelloMessage(name string) string {
	return fmt.Sprintf("Hello World from Go Modules! My name is: %s", name)
}
