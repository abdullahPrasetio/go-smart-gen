package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateProject(destination string) {
	// URL repository
	repo := "https://github.com/abdullahPrasetio/base-go"

	// Nama folder hasil clone
	folderName := "base-go"

	// Clone repository
	if _, err := os.Stat(filepath.Join(folderName, ".git")); !os.IsNotExist(err) {
		// Menghapus folder hasil clone
		err = RemoveFolderAndFile(folderName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	err := exec.Command("git", "clone", repo).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Cek apakah di direktori tujuan sudah ada folder .git
	if _, err := os.Stat(filepath.Join(destination, ".git")); !os.IsNotExist(err) {
		fmt.Println("Folder .git sudah ada di direktori tujuan.")
		return
	}
	// Copy folder hasil clone ke folder tujuan
	err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// Menampilkan loading
		fmt.Print(".")

		rel, err := filepath.Rel(folderName, path)
		if err != nil {
			return err
		}

		dest := filepath.Join(destination, rel)
		destDir := filepath.Dir(dest)
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			err = os.MkdirAll(destDir, 0755)
			if err != nil {
				return err
			}
		}

		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer dst.Close()

		srcBytes, err := ioutil.ReadAll(src)
		if err != nil {
			panic(err)
		}

		_, err = dst.Write(srcBytes)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Menghapus folder hasil clone
	err = RemoveFolderAndFile(folderName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Menghapus folder .git hasil clone
	err = RemoveFolderAndFile(destination + "/.git")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\nSelesai")
}

func RemoveFolderAndFile(folderName string) error {
	// Menghapus folder hasil clone
	err := exec.Command("rm", "-rf", folderName).Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
