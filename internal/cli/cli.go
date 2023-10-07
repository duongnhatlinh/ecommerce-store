package cli

import (
	"ecommercestore/internal/logging"
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Print(`This program runs backend server.
Usage:
Supported arguments:
`)
	flag.PrintDefaults()
	os.Exit(1)
}

func Parse() {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	logging.ConfigureLogger(*env)
	if *env == "prod" {
		logging.SetLogToFile()
	}
}
