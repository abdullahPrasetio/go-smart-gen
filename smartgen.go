package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abdullahPrasetio/go-smart-gen/generator"
)

var (
	make = flag.String("make", "", "Please write make")
)

func main() {
	flag.Parse()
	if *make == "" {
		fmt.Println("Error: Parameter --make harus ditentukan")
		os.Exit(1)
	}
	makeFile := *make
	switch makeFile {
	case "model":
		generator.CreateEntityV2()
	default:
		// fields := []generator.Field{
		// 	{
		// 		Name:    "ID",
		// 		Kind:    "int",
		// 		TagJson: "id",
		// 	},
		// 	{
		// 		Name:    "Name",
		// 		Kind:    "string",
		// 		TagJson: "name",
		// 	},
		// }
		// generator.CreateService("User", fields)
		generator.CreateEntityV2()
	}
}
