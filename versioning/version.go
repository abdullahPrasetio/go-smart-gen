package versioning

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/tcnksm/go-latest"
)

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
	cmd := exec.Command("go", "install", "github.com/abdullahPrasetio/go-smart-gen@"+version)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	fmt.Println("jika versi belum ter update silahkan tambahakan --newversion={v1.x.x}")
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
