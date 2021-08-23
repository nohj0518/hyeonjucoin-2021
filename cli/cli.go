package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/nohj0518/hyeonjucoin-2021/explorer"
	"github.com/nohj0518/hyeonjucoin-2021/rest"
)

func usage(){
	fmt.Printf("Welcome to 현주 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set th PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}      

func Start() {

	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set port the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()
	port_parameter := fmt.Sprintf("127.0.0.1:%d", *port)

	switch *mode {
	case "rest":
		rest.Start(port_parameter)
		// Start rest api
	case "html":
		explorer.Start(port_parameter)
		// Start html explorer
	case "all":
		go rest.Start(port_parameter)
		port_parameter = fmt.Sprintf("127.0.0.1:%d", *port + 1000)
		explorer.Start(port_parameter)
		// Start html all
	default:
		usage()
	}

	fmt.Println(*port, *mode)

}