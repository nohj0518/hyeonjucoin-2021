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

type balanceResponse struct {
	Address string `json:"address"`
	Balance int `json:"balance"`
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type addTxPayload struct{
	To     string
	Amount int
}

func documentation(rw http.ResponseWriter, r *http.Request){
	data := []urlDescription{
		{
			URL: url("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: url("/status"),
			Method: "GET",
			Description: "See the Status of the Blockchain",
		},
		{ 
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
		{ 
			URL: url("/blocks/{hash}"),
			Method: "GET",
			Description: "See a Block",
		},
		{
			URL: url("/balance/{address}"),
			Method: "GET",
			Description: "Get TxOuts for an Address",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain()))
	case "POST":
		blockchain.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound{
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else{
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddlewear(next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func status (rw http.ResponseWriter, r *http.Request){
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchain.BalanceByAddress(address,blockchain.Blockchain())
		json.NewEncoder(rw).Encode(balanceResponse{address,amount})
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.UTxOutsByAddress(address,blockchain.Blockchain())))
	}
}

func mempool(rw http.ResponseWriter, r *http.Request){
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool.Txs))
}

func transactions(rw http.ResponseWriter, r *http.Request){
	var payload addTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		json.NewEncoder(rw).Encode(errorResponse{"Not enough funds"})
	}
	rw.WriteHeader(http.StatusCreated)
}


func Start(aPort string) {
	port = aPort
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddlewear)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET","POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}",balance)
	router.HandleFunc("/mempool",mempool)
	router.HandleFunc("/transactions",transactions).Methods("POST")
	fmt.Printf("Listening on http://%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}