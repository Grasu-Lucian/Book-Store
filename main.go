package main

import (
	"fmt"
	"main/initializers"
)

func init() {
initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello")
}
