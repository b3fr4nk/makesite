package main

import (
	"html/template"
	"os"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content string
}

func main() {
	// read input
	textFilePath := "first-post.txt"
	textFileName := "new"

	fileContents, err := os.ReadFile(textFilePath)

	if err!= nil {
		panic(err)
	}

	//create page struct
	page := Page{
		TextFilePath: textFilePath,
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
