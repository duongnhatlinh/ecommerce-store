package main

import (
	"ecommercestore/internal/app"
	"ecommercestore/internal/cli"
	"ecommercestore/internal/conf"
)

func main() {
	cli.Parse()
	app.Start(conf.NewConfig())
}
