package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page because a wiki is a series of interconneted pages each having following variables
// This defines how each page data is going to be stored in memory.
type Page struct {
	// It is going to be the file name and will also work as the title in the page.
	Title string
	// Body here is a slice of byte. Slice can have any memory amount limited by the os memory.
	// It could be string as well but the library io except it to be in byte.
	Body []byte
}

// save to save each page in the memory
// returns nil if all goes well
// this is a method(not function) and takes a Page object as its recevier.
func (p Page) save() error {
	filename := p.Title + ".txt"
	// 0600 is that the file should bt created with the read-write permissions for the current user only.
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// func loadPage(title string) *Page {
// 	filename := title + ".txt"
// 	body, _ := ioutil.ReadFile(filename)
// 	return &Page{Title: title, Body: body}
// }

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		// Can't return nil here if I don't use pointer return type.
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// http.ResponseWriter sends the data to http client make the request to the server.
// http.Request the resource that our http client wants to access.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// p1 := Page{Title: "page", Body: []byte("This is my page.")}
	// p1.save()
	// p2, _ := loadPage("page")
	// fmt.Println(string(p2.Body))
}
