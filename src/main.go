package main

import (
	"fmt"
	"github.com/tyler-sommer/stick"
	"os"
)

const Template string = "Hello, {{ name }}!\n"

func main() {
	fmt.Println("---Starting---")
	defer fmt.Println("---Ending---")

	env := stick.New(nil)
	env.Execute(Template, os.Stdout, map[string]stick.Value{"name": "Giancarlos"})
}
