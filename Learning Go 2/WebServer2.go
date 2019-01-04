package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"strconv"
	"encoding/json"
)

type Record struct {
	Id    int                `json:"id"`
	Name  string             `json:"name"`
	Phone string             `json:"phone"`
	Class string             `json:"class"`
}

var records = make(map[string]Record)
var recordId int = 1

func main() {

	//sample record
	rec1 := Record{1, "Aman", "9828828228", "A"}

	records[strconv.Itoa(recordId)] = rec1

	r := mux.NewRouter()
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr: ":8000",
		Handler: r,
	}
	log.Println("Server Running...")
	server.ListenAndServe()
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var data []Record
	for _, v := range records {
		data = append(data, v)
	}

	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)

}

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var newData Record
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		fmt.Fprintln(w, "error :")
		return
	}

	recordId++
	fmt.Println(recordId)

	newData.Id = recordId
	records[strconv.Itoa(recordId)] = newData
	jsonData, _ := json.Marshal(newData)

	fmt.Fprintln(w, "success POST : ", string(jsonData))

}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	editId := vars["id"]

	var toBeUpdatedData Record
	err := json.NewDecoder(r.Body).Decode(&toBeUpdatedData)
	if err != nil {
		fmt.Fprintln(w, "error :")
		return
	}

	if _ , ok :=records[editId]; ok{
		delete(records , editId)
		id,_:=strconv.Atoi(editId)
		toBeUpdatedData.Id = id
		records[editId] = toBeUpdatedData
		fmt.Fprintln(w, "PUT Success " , records[editId])

	}else {
		fmt.Fprintln(w, "Enter valid Id for put !")
	}

}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	editId := vars["id"]

	if _ , ok :=records[editId]; ok{
		fmt.Fprintln(w, "This record is DELETED " , records[editId])
		delete(records , editId)

	}else {
		fmt.Fprintln(w, "Enter valid Id for delete !")
	}

}