package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/abdullahPrasetio/go-smart-gen/generator"
	"github.com/tcnksm/go-latest"
)

const TagVersion = "v1.0.0"

var (
	make       = flag.String("make", "", "Please write make")
	folder     = flag.String("folder", "", "Please write folder")
	version    = flag.Bool("version", false, "Please write version")
	update     = flag.Bool("update", false, "Please write update")
	newVersion = flag.String("newversion", "", "Please write update")
)

func main() {
	flag.Parse()
	// if *make == "" {
	// 	fmt.Println("Error: Parameter --make harus ditentukan")
	// 	os.Exit(1)
	// }
	makeFile := *make

	if *version {
		// fmt.Println(*version)
		// GetCurrentVersion()
		CheckVersion()
	}
	if *update {
		newVersioning := *newVersion
		UpdateVersion(newVersioning)
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
		// generator.CreateEntityV2()
		fmt.Println("Welcome to smartgen")
	}
}

func CheckVersion() {
	githubTag := &latest.GithubTag{
		Owner:      "abdullahPrasetio",
		Repository: "go-smart-gen",
	}
	currentVersion := GetCurrentVersion()
	// Periksa apakah versi terbaru tersedia.
	res, err := latest.Check(githubTag, currentVersion)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newVersion := "v" + res.Current
	// fmt.Println(res)

	// Jika versi terbaru tersedia, berikan tahu pengguna.
	if res.Outdated {
		fmt.Printf("Versi terbaru %s tersedia! (Anda menggunakan versi %s)\n", newVersion, currentVersion)
		fmt.Println("Silahkan jalankan program dan tambahkan --update untuk melakukan update generator")
	} else {
		fmt.Printf("Anda menggunakan versi terbaru %s\n", newVersion)
	}
}

func UpdateVersion(newVersion string) {
	version := "latest"
	if newVersion != "" {
		version = newVersion
	}
	fmt.Println(version)
	cmd := exec.Command("go", "install", "github.com/abdullahPrasetio/go-smart-gen@"+version)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fmt.Println("sukses update versi")
	return
}

func GetCurrentVersion() string {
	// nama modul yang ingin diperiksa
	moduleName := "github.com/abdullahPrasetio/go-smart-gen"

	// jalankan perintah 'go list -m' dengan nama modul
	cmd := exec.Command("go", "list", "-m", "-versions", moduleName)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// baca output sebagai string
	outputString := string(output)

	// parsing output untuk mendapatkan versi
	fields := strings.Fields(outputString)
	// fmt.Println(outputString)
	version := fields[len(fields)-1]

	// tampilkan versi
	// fmt.Println("Versi", moduleName, "yang terinstal:", version)
	return version

}
