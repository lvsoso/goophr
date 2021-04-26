package main


import (
	"fmt"
	"log"
	"net/http"

	booksHandler "restServer/books-handler"
)


func main(){
	books := booksHandler.GetBooks()
	log.Println(fmt.Sprintf("%+v", books))

	actionCh := make(chan booksHandler.Action)

	go booksHandler.StartBooksManager(books, actionCh)

	http.HandleFunc("/api/books/", booksHandler.MakeHandler(booksHandler.BookHandler, "/api/books/", actionCh))

	log.Println("Starting Server at port 8080 ...")
	http.ListenAndServe(":8080", nil)
}
