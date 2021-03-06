/*
 * Books API
 *
 * This web service offers information on books
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strings"
)

var authors = []Author{
	{
		AuthorId:    "Author1",
		Name:        "Pedro S. Gonzalez",
		Nationality: "US",
		Birth:       "1970-02-11",
		Genere:      "Sports",
		Books:       nil,
	},
	{
		AuthorId:    "Author2",
		Name:        "Andrew S. Tanenbaum",
		Nationality: "US",
		Birth:       "1970-02-11",
		Genere:      "Technology",
		Books:       nil,
	},
}

var publishers = []Publisher{
	{
		PublisherId: "Publisher1",
		Name:        "John Wiley & Sons",
		Country:     "US",
		Founded:     "2012",
		Genere:      "Technology",
		Books:       nil,
	},
	{
		PublisherId: "Publisher2",
		Name:        "Willy S. Figueres",
		Country:     "US",
		Founded:     "2010",
		Genere:      "Arts",
		Books:       b},
}

var b *Book
var testBook = Book{
	BookId: "Book99", Title: "Computer Networks", Edition: "5th",
	Copyright: "2010", Language: "ENGLISH", Pages: "960",
	Authors: []Author{
		{
			AuthorId:    "Author1",
			Name:        "Pedro S. Gonzalez",
			Nationality: "US",
			Birth:       "1970-02-11",
			Genere:      "Sports",
			Books:       nil,
		}},
	Publishers: []Publisher{
		{
			PublisherId: "Publisher2",
			Name:        "Willy S. Figueres",
			Country:     "US",
			Founded:     "2010",
			Genere:      "Arts",
			Books:       b}},
}
var books = []Book{
	{
		BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
		Copyright: "2012", Language: "ENGLISH", Pages: "976",
		Authors:
		[]Author{
			{
				AuthorId:    "Author2",
				Name:        "Andrew S. Tanenbaum",
				Nationality: "US",
				Birth:       "1970-02-11",
				Genere:      "Technology",
				Books:       nil,
			},
		},
		Publishers: []Publisher{
			{
				PublisherId: "Publisher1",
				Name:        "John Wiley & Sons",
				Country:     "US",
				Founded:     "2012",
				Genere:      "Technology",
				Books:       nil,
			},
		},
	},
	{
		BookId: "Book3", Title: "Computer Networks", Edition: "5th",
		Copyright: "2010", Language: "ENGLISH", Pages: "960",
		Authors: []Author{
			{
				AuthorId:    "Author2",
				Name:        "Andrew S. Tanenbaum",
				Nationality: "US",
				Birth:       "1970-02-11",
				Genere:      "Technology",
				Books:       nil,
			}},
		Publishers: []Publisher{
			{
				PublisherId: "Publisher2",
				Name:        "Willy S. Figueres",
				Country:     "US",
				Founded:     "2010",
				Genere:      "Arts",
				Books:       b}},
	},
}

func findBook(x string) int {
	for i, book := range books {
		if x == book.BookId {
			return i
		}
	}
	return -1
}

func findAuthor(x string) int {
	for i, element := range authors {
		if x == element.AuthorId {
			return i
		}
	}
	return -1
}

func findPublisher(x string) int {
	for i, element := range publishers {
		if x == element.PublisherId {
			return i
		}
	}
	return -1
}

func AuthorsAuthorIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findAuthor(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if i == -1 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		authors = append(authors[:i], authors[i+1:]...)
		w.WriteHeader(http.StatusOK)
	}
}

func AuthorsAuthorIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findAuthor(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(authors[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func AuthorsGet(w http.ResponseWriter, r *http.Request) {
	dataJson, _ := json.Marshal(authors)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func AuthorsBooksGet(w http.ResponseWriter, r *http.Request) {
	authorBooks := []Book {}
	prev := path.Dir(r.URL.Path)
	urll, _ := url.Parse(prev)
	authorId := path.Base(urll.Path)

	i := findAuthor(authorId)
	if i != -1 {
		author := authors[i]
		//dataJson, _ := json.Marshal(author.Books)
		for _, book := range books {
			for _, bookAuthor := range book.Authors {
				if strings.Compare(author.AuthorId, bookAuthor.AuthorId) == 0 {
					authorBooks = append(authorBooks,book)
				}
				//author.Books = &book
				//authors = append(authors, author)
			}
		}
		dataJson, _ := json.Marshal(authorBooks)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(dataJson)
		w.WriteHeader(http.StatusOK)
	}
	dataJson2, _ := json.Marshal(i)
	w.Write(dataJson2)
}

func AuthorsAuthorIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func AuthorsPost(w http.ResponseWriter, r *http.Request) {
	var author Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authors = append(authors, author)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBook(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if i == -1 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		books = append(books[:i], books[i+1:]...)
		w.WriteHeader(http.StatusOK)
	}
}

func BooksBookIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBook(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func BooksGet(w http.ResponseWriter, r *http.Request) {
	dataJson, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksPost(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PublishersPost(w http.ResponseWriter, r *http.Request) {
	var publisher Publisher
	err := json.NewDecoder(r.Body).Decode(&publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	publishers = append(publishers, publisher)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findPublisher(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if i == -1 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		publishers = append(publishers[:i], publishers[i+1:]...)
		w.WriteHeader(http.StatusOK)
	}
}

func PublishersPublisherIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findPublisher(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(publishers[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func PublishersGet(w http.ResponseWriter, r *http.Request) {
	dataJson, _ := json.Marshal(publishers)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func PublishersBooksGet(w http.ResponseWriter, r *http.Request) {
	publisherBooks := []Book {}
	prev := path.Dir(r.URL.Path)
	urll, _ := url.Parse(prev)
	publisherId := path.Base(urll.Path)
	i := findPublisher(publisherId)
	if i != -1 {
		publisher := publishers[i]
		//dataJson, _ := json.Marshal(publisher.Books)
		for _, book := range books {
			for _, bookPublisher := range book.Publishers {
				if strings.Compare(publisher.PublisherId, bookPublisher.PublisherId) == 0 {
					publisherBooks = append(publisherBooks,book)
				}
				//author.Books = &book
				//authors = append(authors, author)
			}
		}
		dataJson, _ := json.Marshal(publisherBooks)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(dataJson)
		w.WriteHeader(http.StatusOK)
	}
}

func PublishersPublisherIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
