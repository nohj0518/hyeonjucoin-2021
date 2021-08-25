package main

import (
	"github.com/nohj0518/hyeonjucoin-2021/cli"
	"github.com/nohj0518/hyeonjucoin-2021/db"
)

func main() {
	defer db.Close()
	cli.Start()
}