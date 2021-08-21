package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nohj0518/hyeonjucoin-2021/blockchain"
	"github.com/nohj0518/hyeonjucoin-2021/utils"
)

//const port string = "127.0.0.1:3000"// ":4000"
var port string

type url string

func (u url) MarshalText() ([]byte, error){
	url := fmt.Sprintf("http://%s%s",port,u)
	return []byte(url), nil
}
          
type urlDescription struct{
	URL         url `json:"url"`  
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
	
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request){
	data := []urlDescription{
		{
			URL: url("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{ 
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
		{ 
			URL: url("/blocks/{id}"),
			Method: "GET",
			Description: "See a Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
}

func Start(aPort string) {
	router := mux.NewRouter()
	port = aPort
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET","POST")
	router.HandleFunc("/blocks/{id:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}