package booksHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func StartBooksManager(books map[string]bookResource, actionCh <-chan Action){
	newID := len(books)

	for {
		select {
		case act := <-actionCh:
			switch act.Type {
			case "GET":
				actOnGET(books, act)
			case "POST":
				newID++
				newBookID := fmt.Sprintf("%d", newID)
				actOnPOST(books, act, newBookID)
			case "PUT":
				actOnPUT(books, act)
			case "DELETE":
				actOnDELETE(books, act)
			}
		}
	}
}


func BookHandler(w http.ResponseWriter, r *http.Request, id string, method string, actionCh chan<- Action){

	isGet := method == "GET"
	idIsSetForPost := method == "POST" && id != ""
        isPutOrPost := method == "PUT" || method == "POST"
        idIsSetForDelPut := (method == "DELETE" || method == "PUT") && id != ""

	if !isGet && !(idIsSetForPost || idIsSetForDelPut || isPutOrPost){
		writeError(w, http.StatusMethodNotAllowed)
                return
	}


        respCh := make(chan response)
        act := Action{
                Id:      id,
                Type:    method,
                RetChan: respCh,
        }

	if isPutOrPost {
		var reqPayload requestPayload
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err := json.Unmarshal(body, &reqPayload); err != nil {
			writeError(w, http.StatusBadRequest)
			return
		}

		act.Payload = reqPayload
	}


	actionCh <- act

	var resp response
	if resp = <-respCh; resp.StatusCode > http.StatusCreated {
			writeError(w, resp.StatusCode)
			return
	}

	if method == "DELETE" {
			log.Println(fmt.Sprintf("Resource ID %s deleted: %+v", id, resp.Books))
			resp = response{
					StatusCode: http.StatusOK,
					Books:      []bookResource{},
			}
	}

	writeResponse(w, resp)

}
