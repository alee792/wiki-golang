package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func viewHandler(w http:ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</hi><div>%s</div>", p.Title, p.Body)
}


func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
