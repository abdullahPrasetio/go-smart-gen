package generator

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	asset "github.com/abdullahPrasetio/go-smart-gen/template"
)

type Field struct {
	Name       string
	Kind       string
	TagJson    string
	TagBinding string
}

func CreateEntityV1() {
	fmt.Print("Enter model name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	modelName := scanner.Text()
	allKind := ""
	var fields []Field

	for {
		fmt.Print("Enter field name (or enter to stop): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fieldName := scanner.Text()

		if fieldName == "" {
			break
		}

		fmt.Print("Enter field type: ")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fieldType := scanner.Text()
		allKind += fieldType

		fmt.Print("Enter field tag json: ")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fieldTag := scanner.Text()

		fmt.Print("Enter field validate: ")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fieldTagBinding := scanner.Text()

		if fieldTag == "" {
			fieldTag = fieldName
		}
		fmt.Println(fieldTag)
		fmt.Println(fieldName)

		fields = append(fields, Field{Name: fieldName, Kind: fieldType, TagJson: fieldTag, TagBinding: fieldTagBinding})
	}
	// Create Folder
	nameLower := strings.ToLower(modelName)
	folder := "./models/" + nameLower + "/"
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Println("Error membuat folder:", err)
		os.Exit(1)
	}
	fullpath := folder + "entity" + ".go"
	f, err := os.Create(fullpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Fprintln(f, "package "+strings.ToLower(modelName))
	fmt.Fprintln(f)
	if strings.Contains(allKind, "sql") {

		fmt.Fprintln(f, "import \"database/sql\"")
	}
	fmt.Fprintln(f)

	fmt.Fprintln(f, "type "+modelName+" struct {")
	for _, field := range fields {
		fmt.Fprint(f, "\t"+field.Name+"\t"+field.Kind)
		if field.TagJson != "" || field.TagBinding != "" {
			fmt.Fprint(f, " `")
			if field.TagJson != "" {
				fmt.Fprint(f, ""+"json:\""+field.TagJson+"\"")
				if field.TagBinding != "" {
					fmt.Fprint(f, " ")
				}
			}
			if field.TagBinding != "" {
				fmt.Fprint(f, ""+"binding:\""+field.TagBinding+"\"")
			}
			fmt.Fprint(f, "`")
		}
		fmt.Fprint(f, "\n")
	}
	fmt.Fprintln(f, "}")
	exec.Command("gofmt -w " + fullpath).Output()
}

func CreateEntityV2() {
	fmt.Print("Enter model name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	modelName := scanner.Text()
	allKind := ""
	var fields []Field

	for {
		fmt.Print("Enter field name (or enter to stop): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fieldName := scanner.Text()

		if fieldName == "" {
			break
		}
		changeFieldType := true
		fieldType := ""
		for changeFieldType {
			fmt.Print("Enter field type: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			fieldType = scanner.Text()
			allKind += fieldType
			if fieldType != "" {
				changeFieldType = false
			}
		}

		fmt.Print("Enter field tag json: ")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fieldTag := scanner.Text()

		fmt.Print("Enter field validate: ")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fieldTagBinding := scanner.Text()
		if fieldTag == "" {
			fieldTag = fieldName
		}

		fields = append(fields, Field{Name: fieldName, Kind: fieldType, TagJson: fieldTag, TagBinding: fieldTagBinding})
	}
	// Create Folder
	nameLower := strings.ToLower(modelName)
	folder := "./models/" + nameLower + "/"
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Println("Error membuat folder:", err)
		os.Exit(1)
	}
	fullpath := folder + "entity" + ".go"
	err = ChangeModelTemplateVariables(fullpath, modelName)
	if err != nil {
		panic("change model template variables error")
	}
	err = AddNewLineInModelTemplate(fullpath, fields, modelName, allKind)
	if err != nil {
		panic("change model template variables error")
	}
	err = CreateRepository(modelName, fields)
	if err != nil {
		panic("create repository error" + err.Error())
	}
	err = CreateService(modelName, fields)
	if err != nil {
		panic("create service error" + err.Error())
	}

}
func ChangeModelTemplateVariables(filePath string, modelName string) error {
	// fmt.Println("assets", asset.Assets.Files["/template/entity.txt"].Data)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error membuat file:", err)
		os.Exit(1)
	}
	defer file.Close()
	// read template file
	// tmplBytes, err := ioutil.ReadFile("/entity.txt")
	// if err != nil {
	// 	fmt.Println("Error reading template file:", err)
	// 	return err
	// }
	// read from binary
	tmplBytes := asset.Assets.Files["/template/entity.txt"].Data
	tpl, err := template.New("file").Parse(string(tmplBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	// Menuliskan isi file dari template
	err = tpl.Execute(file, struct {
		Package string
		Name    string
	}{
		Package: strings.ToLower(modelName),
		Name:    strings.Title(modelName),
	})
	if err != nil {
		fmt.Println("Error menuliskan template:", err)
		os.Exit(1)
	}
	return nil
}

func AddNewLineInModelTemplate(fullPath string, fields []Field, modelName string, allKind string) error {
	//Membuka file template.txt
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	// Baca file dan cari dua komentar yang ditentukan
	scannerTemplate := bufio.NewScanner(file)
	var lines []string
	for scannerTemplate.Scan() {
		lines = append(lines, scannerTemplate.Text())
	}
	// Mencari baris yang sesuai dengan komentar start dan end menggunakan regex
	startCommentRegex := regexp.MustCompile(`// Entity start`)
	endCommentRegex := regexp.MustCompile(`// Entity end`)
	startLineIndex := -1
	endLineIndex := -1
	for i, line := range lines {
		if startCommentRegex.MatchString(line) {
			startLineIndex = i
		}
		if endCommentRegex.MatchString(line) {
			endLineIndex = i
		}
		if startLineIndex != -1 && endLineIndex != -1 {
			break
		}
	}

	//Tulis kembali ke file baru
	newFile, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Error creating new file:", err)
		return err
	}
	defer newFile.Close()

	// Menuliskan isi file template ke file baru, serta menambahkan line baru di antara komentar
	w := bufio.NewWriter(newFile)
	for i, line := range lines {
		if i == startLineIndex+1 {
			// fmt.Fprintln(w, "Tambahkan line baru disini")
			if strings.Contains(allKind, "sql") {

				fmt.Fprintln(w, "import \"database/sql\"")
			}
			fmt.Fprintln(w)

			fmt.Fprintln(w, "type "+modelName+" struct {")
			for _, field := range fields {
				fmt.Fprint(w, "\t"+field.Name+"\t"+field.Kind)
				if field.TagJson != "" || field.TagBinding != "" {
					fmt.Fprint(w, " `")
					if field.TagJson != "" {
						fmt.Fprint(w, ""+"json:\""+field.TagJson+"\"")
						if field.TagBinding != "" {
							fmt.Fprint(w, " ")
						}
					}
					if field.TagBinding != "" {
						fmt.Fprint(w, ""+"binding:\""+field.TagBinding+"\"")
					}
					fmt.Fprint(w, "`")
				}
				fmt.Fprint(w, "\n")
			}
			fmt.Fprintln(w, "}")
		}

		fmt.Fprintln(w, line)
	}

	w.Flush()
	return nil
}
