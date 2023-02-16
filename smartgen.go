package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abdullahPrasetio/go-smart-gen/generator"
	"github.com/abdullahPrasetio/go-smart-gen/versioning"
)

const TagVersion = "v1.0.0"

var (
	make       = flag.String("make", "", "Create model,project dll")
	folder     = flag.String("folder", "", "Create folder for create project")
	version    = flag.Bool("version", false, "Check Version")
	update     = flag.Bool("update", false, "Update version generator")
	newVersion = flag.String("newversion", "", "Change a new version")
)

func main() {
	// Membuat flag
	helpFlag := flag.Bool("h", false, "print help message")

	// Mengubah penggunaan flag
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Parse flag
	flag.Parse()

	// Cek apakah help flag di-set
	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}
	// if *make == "" {
	// 	fmt.Println("Error: Parameter --make harus ditentukan")
	// 	os.Exit(1)
	// }
	makeFile := *make

	if *version {
		// fmt.Println(*version)
		// GetCurrentVersion()
		versioning.CheckVersion()
	}
	if *update {
		newVersioning := *newVersion
		versioning.UpdateVersion(newVersioning)
	}

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
		// fmt.Println("Welcome to smartgen")
	}

	fmt.Println("Untuk mendapatkan bantuan tambah tag -h or -help")
	fmt.Println("Welcome to smartgen")
	fmt.Println("Jika anda merasa aplikasi ini berguna silahkan kunjungi beberapa project saya kritik dan saran bisa di tambahkan di issue")
	fmt.Println("My Project")
	fmt.Println("- https://github.com/abdullahPrasetio/go-smart-gen")
	fmt.Println("- https://github.com/abdullahPrasetio/base-go")
	fmt.Println("- https://github.com/abdullahPrasetio/validation-formatter")
}
