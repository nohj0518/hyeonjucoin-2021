package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/nohj0518/hyeonjucoin-2021/explorer"
	"github.com/nohj0518/hyeonjucoin-2021/rest"
)

func usage() {
	fmt.Printf("Welcome to 현주 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set th PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set port the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
		// Start rest api
	case "html":
		explorer.Start(*port)
		// Start html explorer
	case "all":
		go rest.Start(*port)
		*port += 1000
		explorer.Start(*port)
		// Start html all
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
