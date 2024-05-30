package main

import (
	"FinalTaskAPI/router"
	"fmt"
)

func main() {
	fmt.Print("Hello, World!")
	r := router.StartApp()

	r.Run(":8080")
}
