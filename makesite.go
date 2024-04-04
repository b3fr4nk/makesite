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
	// read open directory and return all text files
	files, err := os.ReadDir(path)
	txtFiles := []string{}
	if err != nil {
		panic(err)
	}

	for _, file := range(files) {
		if file.IsDir() {
			subDirTxtFiles, err := readDir(path+"/"+file.Name())
			if err != nil {
				panic(err)
			}
			for i := range(subDirTxtFiles) {
				subDirTxtFiles[i] = file.Name() + "/" + subDirTxtFiles[i]
			}
			txtFiles = append(txtFiles, subDirTxtFiles...)
		}
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
	return nil
} 

func main() {
	count := 0

	const (
        Reset  = "\033[0m"
        Red    = "\033[31m"
        Green  = "\033[32m"
        Yellow = "\033[33m"
        Blue   = "\033[34m"
        Purple = "\033[35m"
        Cyan   = "\033[36m"
        White  = "\033[37m"
    )

	const (
        Revert = "\033[0m"
        Bold  = "\033[1m"
    )

	// read input
	textFilePathPtr := flag.String("file", "", "path of .txt file in the current directory.")
	textFileDirPtr := flag.String("dir", "~/Dev/makesite/", "path to text files for static site generation")
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
			err = createPage(path, strings.Trim(path, ".txt"), string(fileContents))
			if err != nil {
				panic(err)
			}
			count++
		}
	}

	fmt.Println(Green + Bold + "Success! " + Revert + White + "Generated", count, "Pages")
	
}
