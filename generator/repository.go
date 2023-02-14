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

func CreateRepository(modelName string, fields []Field) error {
	filePath := "./models/" + modelName + "/repository.go"
	fieldString := ""
	countField := len(fields)
	fieldScan := ""
	for index, field := range fields {
		fieldScan += "&result." + field.Name
		if field.TagJson != "" {
			fieldString += "\"" + field.TagJson + "\""
		} else {
			fieldString += "\"" + field.Name + "\""
		}
		if index+1 != countField {
			fieldString += ","
			fieldScan += ","
		}
	}

	fieldQuery := "[]string{" + fieldString + "}"
	err := ChangeRepositoryTemplateVariables(filePath, modelName, fieldQuery, fieldScan)

	for _, field := range fields {
		err = AddLineNewFunctionWithWhere(filePath, modelName, fieldQuery, fieldScan, field)
	}
	return err
}

func ChangeRepositoryTemplateVariables(filePath string, modelName string, fieldQuery string, fieldScan string) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error membuat file:", err)
		os.Exit(1)
	}
	defer file.Close()
	// read from file
	// tmplBytes, err := ioutil.ReadFile("/repository.txt")
	// if err != nil {
	// 	fmt.Println("Error reading template file:", err)
	// 	return err
	// }

	// Read from binary
	tmplBytes := asset.Assets.Files["/template/repository.txt"].Data
	tpl, err := template.New("file").Parse(string(tmplBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	// Menuliskan isi file dari template
	err = tpl.Execute(file, struct {
		Package    string
		Name       string
		FieldQuery string
		FieldScan  string
		NameTable  string
	}{
		Package:    strings.ToLower(modelName),
		Name:       strings.Title(modelName),
		FieldQuery: fieldQuery,
		FieldScan:  fieldScan,
		NameTable:  strings.ToLower(modelName) + "s",
	})
	if err != nil {
		fmt.Println("Error menuliskan template:", err)
		os.Exit(1)
	}
	return nil
}

func AddLineNewFunctionWithWhere(filePath string, modelName string, fieldQuery string, fieldScan string, field Field) error {
	// Baca file1 dan ubah isinya menggunakan template search
	// tmpl, err := template.ParseFiles("./template/repository_search.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// fieldSearchDatabase := ""
	// if field.TagJson != "" {
	// 	fieldSearchDatabase = field.TagJson
	// } else {
	// 	fieldSearchDatabase = field.Name
	// }

	// Read from binary
	tmplBytes := asset.Assets.Files["/template/repository_search.txt"].Data
	tmpl, err := template.New("file").Parse(string(tmplBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}
	var result string
	var b strings.Builder
	err = tmpl.Execute(&b, struct {
		Package             string
		Name                string
		FieldQuery          string
		FieldScan           string
		FieldSearch         string
		FieldSearchDatabase string
		NameTable           string
	}{
		Package:             strings.ToLower(modelName),
		Name:                strings.Title(modelName),
		FieldQuery:          fieldQuery,
		FieldScan:           fieldScan,
		FieldSearch:         field.Name,
		FieldSearchDatabase: field.TagJson,
		NameTable:           strings.ToLower(modelName) + "s",
	})
	result = b.String()
	if err != nil {
		fmt.Println(err)
		return err
	}
	// read from file
	// tmplInterface, err := template.ParseFiles("./template/repository_search_interface.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	tmplBytesInterface := asset.Assets.Files["/template/repository_search_interface.txt"].Data
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
