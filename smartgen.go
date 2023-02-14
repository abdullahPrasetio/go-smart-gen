package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abdullahPrasetio/go-smart-gen/generator"
)

var (
	make   = flag.String("make", "", "Please write make")
	folder = flag.String("folder", "", "Please write folder")
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
	case "project":
		if *folder == "" {
			fmt.Println("Error: Jika anda ingin create project Parameter --folder harus ditentukan")
			os.Exit(1)
		}
		generator.CreateProject(*folder)
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
