package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

// This function saves a Page's Body to a text file
// The name of the file is the Page's Title
// nil is returned upon sucess
// error is returned otherwise
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// Third arg refers to RW permissions
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// This function takes a title string and loads a file
// Looks like you define the return
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		// Return no Page, and the error
		return nil, err
	}
	// Returns a new Page literal and nil for no error
	return &Page{Title: title, Body: body}, nil
}

// Create a test page and save it
// Load it and dump any errors
// Print it
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}