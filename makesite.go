package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content string
}

func readDir (path string) ([]string, error) {
	files, err := os.ReadDir(path)
	txtFiles := []string{}
	if err != nil {
		panic(err)
	}

	for _, file := range(files) {
		_, path, isCut:= strings.Cut(file.Name(), ".")
		if isCut {
			if path == "txt" {
				txtFiles = append(txtFiles, file.Name())
			}
		}
	}

	return txtFiles, nil
}

func createPage(filePath, fileName, Contents string) (error) {
	//create page struct
	page := Page{
		TextFilePath: filePath,
		TextFileName: fileName,
		HTMLPagePath: fileName+".html",
		Content: Contents,
	}

	// Create template
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		return err
	}

	//generate html
	err = t.Execute(newFile, page)
	if err != nil {
		return err
	}

	fmt.Println("done")

	return nil
} 

func main() {
	// read input
	textFilePathPtr := flag.String("file", "", "path of .txt file in the current directory.")
	textFileDirPtr := flag.String("dir", "./", "path to text files for static site generation")
	flag.Parse()	
	
	if *textFilePathPtr != "" {
		textFileName := strings.Trim(*textFilePathPtr, ".txt")
		fileContents, err := os.ReadFile(*textFilePathPtr)

		if err!= nil {
			panic(err)
		}

		createPage(*textFilePathPtr, textFileName, string(fileContents))

		
	} else {
		files, err := readDir(*textFileDirPtr)
		if err != nil {
			panic(err)
		}
		
		for _, file := range(files) {
			path := *textFileDirPtr + "/" + file
			fileContents, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			fmt.Println(file + " htmling")
			err = createPage(path, strings.Trim(path, ".txt"), string(fileContents))
			if err != nil {
				panic(err)
			}
		}
	}

	
	
}
