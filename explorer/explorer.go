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
	templateDir string = "explorer/templates/"
)


type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	
	data := homeData{"Home", nil }
	templates.ExecuteTemplate(rw,"home", data)
}

func add(rw http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.Blockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
	
}

func Start(port string) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}