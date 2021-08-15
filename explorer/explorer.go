package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/nohj0518/hyeonjucoin-2021/blockchain"
)
var templates *template.Template

const (
	port string = ":4000"
	templateDir string = "explorer/templates/"
)


type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw,"home", data)
}

func add(rw http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
	
}

func Start() {
	fmt.Println("5.4 Adding A Block !")
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}