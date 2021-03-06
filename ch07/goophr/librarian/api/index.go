package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// path /index 相关逻辑

type tPayload struct {
	Token  string `json:"token"`
	Title  string `json:"title"`
	DocID  string `json:"doc_id"`
	LIndex int    `json:"line_index"`
	Index  int    `json:"token_index"`
}

type tIndex struct {
	Index  int
	LIndex int
}

func (ti tIndex) String() string {
	return fmt.Sprintf("i: %d, li: %d", ti.Index, ti.LIndex)
}


type tIndices []tIndex

type document struct {
	Count   int
	DocID   string
	Title   string
	Indices map[int]tIndices
}

func (d *document) String() string {
	str := fmt.Sprintf("%s (%s): %d\n", d.Title, d.DocID, d.Count)
	var buffer bytes.Buffer

	for lin, tis := range d.Indices {
		var lBuffer bytes.Buffer
		for _, ti := range tis {
			lBuffer.WriteString(fmt.Sprintf("%s ", ti.String()))
		}
		buffer.WriteString(fmt.Sprintf("@%d -> %s\n", lin, lBuffer.String()))
	}
	return str + buffer.String()
}

type documentCatalog map[string]*document

func (dc *documentCatalog) String() string {
	return fmt.Sprintf("%#v", dc)
}

type tCatalog map[string]documentCatalog

func (tc *tCatalog) String() string {
	return fmt.Sprintf("%#v", tc)
}

type tcCallback struct {
	Token string
	Ch    chan tcMsg
}

type tcMsg struct {
	Token string
	DC    documentCatalog
}

var pProcessCh chan tPayload

var tcGet chan tcCallback

func StartIndexSystem() {
	pProcessCh = make(chan tPayload, 100)
	tcGet = make(chan tcCallback, 20)
	go tIndexer(pProcessCh, tcGet)
}

func tIndexer(ch chan tPayload, callback chan tcCallback)  {
	store := tCatalog{}
	for {
		select {
		case msg := <- callback:
			dc := store[msg.Token]
			msg.Ch <- tcMsg{
				Token: msg.Token,
				DC:    dc,
			}
		case pd := <-ch:
			dc, exists := store[pd.Token]
			if !exists{
				dc = documentCatalog{}
				store[pd.Token] = dc
			}

			doc, exists := dc[pd.DocID]
			if !exists {
				doc = &document{
					DocID:   pd.DocID,
					Title:   pd.Title,
					Indices: map[int]tIndices{},
				}
				dc[pd.DocID] = doc
			}

			tin := tIndex{
				Index: pd.Index,
				LIndex: pd.LIndex,
			}

			doc.Indices[tin.LIndex] = append(doc.Indices[tin.LIndex], tin)
			doc.Count ++
		}
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"code": 405, "msg": "Method Not Allowed."}`))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var tp tPayload
	decoder.Decode(&tp)
	log.Printf("Token received%#v\n", tp)

	pProcessCh <- tp
	w.Write([]byte(`{"code": 200, "msg": "Tokens are being added to index."}`))
}