package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	asset "github.com/abdullahPrasetio/go-smart-gen/template"
)

func CreateService(modelName string, fields []Field) error {
	filePath := "./models/" + modelName + "/service.go"
	err := ChangeServiceTemplateVariables(filePath, modelName)
	fmt.Println(err)
	for _, field := range fields {
		err = AddLineServiceNewFunctionWithWhere(filePath, modelName, field)
		fmt.Println(err)
	}
	return err
}

func ChangeServiceTemplateVariables(filePath string, modelName string) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error membuat file:", err)
		os.Exit(1)
	}
	defer file.Close()
	// Read file
	// tmplBytes, err := ioutil.ReadFile("/service.txt")
	// if err != nil {
	// 	fmt.Println("Error reading template file:", err)
	// 	return err
	// }

	tmplBytes := asset.Assets.Files["/template/service.txt"].Data
	tpl, err := template.New("file").Parse(string(tmplBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	// Menuliskan isi file dari template
	err = tpl.Execute(file, struct {
		Package   string
		Name      string
		NameTable string
	}{
		Package:   strings.ToLower(modelName),
		Name:      strings.Title(modelName),
		NameTable: strings.ToLower(modelName) + "s",
	})
	if err != nil {
		fmt.Println("Error menuliskan template:", err)
		os.Exit(1)
	}
	return nil
}

func AddLineServiceNewFunctionWithWhere(filePath string, modelName string, field Field) error {
	// Baca file1 dan ubah isinya menggunakan template search
	// tmpl, err := template.ParseFiles("./template/service_search.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// Read from binary
	tmplBytes := asset.Assets.Files["/template/service_search.txt"].Data
	tmpl, err := template.New("file").Parse(string(tmplBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}
	var result string
	var b strings.Builder
	err = tmpl.Execute(&b, struct {
		Name        string
		FieldSearch string
	}{
		Name:        strings.Title(modelName),
		FieldSearch: field.Name,
	})
	result = b.String()
	if err != nil {
		fmt.Println(err)
		return err
	}
	// read from file
	// tmplInterface, err := template.ParseFiles("/service_search_interface.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	tmplBytesInterface := asset.Assets.Files["/template/service_search_interface.txt"].Data
	tmplInterface, err := template.New("file").Parse(string(tmplBytesInterface))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	var resultInterface string
	var stringInterface strings.Builder
	err = tmplInterface.Execute(&stringInterface, struct {
		Name        string
		FieldSearch string
	}{
		Name:        strings.Title(modelName),
		FieldSearch: field.Name,
	})
	resultInterface = stringInterface.String()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Baca file2 dan cari komentar // Entity start dan // Entity end
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	startCommentInterface := regexp.MustCompile(`// Add Function interface`)
	endCommentInterface := regexp.MustCompile(`// End Function interface`)

	scanner := bufio.NewScanner(file)
	var lines []string
	//inBlockInterface := false
	startLineIndex := -1
	endLineIndex := -1
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i, line := range lines {
		if startCommentInterface.MatchString(line) {
			startLineIndex = i
		}
		if endCommentInterface.MatchString(line) {
			endLineIndex = i
		}
		if startLineIndex != -1 && endLineIndex != -1 {
			break
		}
	}
	newLines := []string{}
	for i, line := range lines {
		if i == startLineIndex+1 {
			newLines = append(newLines, resultInterface)
		}
		newLines = append(newLines, line)
	}
	startCommentFunction := regexp.MustCompile(`// Add New Function`)
	endCommentFunction := regexp.MustCompile(`// End New Function`)
	startNewLineIndex := -1
	endNewLineIndex := -1
	for i, newLine := range newLines {
		if startCommentFunction.MatchString(newLine) {
			startNewLineIndex = i
		}
		if endCommentFunction.MatchString(newLine) {
			endNewLineIndex = i
		}
		if startNewLineIndex != -1 && endNewLineIndex != -1 {
			break
		}
	}
	var newLinesFunction []string
	for i, line := range newLines {
		if i == startNewLineIndex+1 {
			newLinesFunction = append(newLinesFunction, result)
		}
		newLinesFunction = append(newLinesFunction, line)
	}
	// Tulis result ke danewLinesFunctionam lines

	// Tulis lines kembali ke file2
	err = ioutil.WriteFile(filePath, []byte(strings.Join(newLinesFunction, "\n\n")), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
