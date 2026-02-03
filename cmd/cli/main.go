package main

import (
	"fmt"

	"github.com/dimka90/taskit/internal"
)

func main() {

	task := internal.SimplifyTask{"dimka", "Cooking", false}
	fmt.Println(task.MarkComplete())
	fmt.Println(task.GetStatus())
}
