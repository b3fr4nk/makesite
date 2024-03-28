package main

import (
	"flag"
	"html/template"
	"os"
	"strings"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content string
}

func main() {
	// read input
	textFilePathPtr := flag.String("file", "new-post.txt", "path of .txt file in the current directory.")
	flag.Parse()
	textFileName := strings.Trim(*textFilePathPtr, ".txt")

	fileContents, err := os.ReadFile(*textFilePathPtr)

	if err!= nil {
		panic(err)
	}

	//create page struct
	page := Page{
		TextFilePath: *textFilePathPtr,
		TextFileName: textFileName,
		HTMLPagePath: textFileName+".html",
		Content: string(fileContents),
	}

	// Create template
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	//generate html
	err = t.Execute(newFile, page)
	if err != nil {
		panic(err)
	}
	
}
