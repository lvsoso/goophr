package booksHandler

import (
	"net/http"
)


func actOnGET(books map[string]bookResource, act Action){
        status := http.StatusNotFound
        bookResult := []bookResource{}

	if act.Id == "" {
		status = http.StatusOK
		for _, book := range books {
			bookResult = append(bookResult, book)
		}
	} else if book, exists := books[act.Id]; exists {
		status = http.StatusOK
		bookResult = []bookResource{book}
	}

	act.RetChan <- response{
                StatusCode: status,
                Books:      bookResult,
	}
}

func actOnDELETE(books map[string]bookResource, act Action) {
        book, exists := books[act.Id]
        delete(books, act.Id)

        if !exists {
                book = bookResource{}
        }

        act.RetChan <- response{
                StatusCode: http.StatusOK,
                Books:      []bookResource{book},
        }
}


func actOnPUT(books map[string]bookResource, act Action) {
        status := http.StatusNotFound
        bookResult := []bookResource{}

	if book, exists := books[act.Id]; exists {
                book.Link = act.Payload.Link
                book.Title = act.Payload.Title
                books[act.Id] = book

                status = http.StatusOK
                bookResult = []bookResource{books[act.Id]}
        }


        act.RetChan <- response{
                StatusCode: status,
                Books:      bookResult,
        }

}


func actOnPOST(books map[string]bookResource, act Action, newID string) {
        books[newID] = bookResource{
                Id:    newID,
                Link:  act.Payload.Link,
                Title: act.Payload.Title,
        }

        act.RetChan <- response{
                StatusCode: http.StatusCreated,
                Books:      []bookResource{books[newID]},
        }
}


