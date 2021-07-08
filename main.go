package main

import (
	"flag"
	"fmt"
	"github.com/RHsyseng/pullsecret-validator-cli/command"
	"os"
)

func main() {

	f := flag.String("f", "", "[Required param] - Pull Secret file in json format. Usage: '-f ./pull-secret.json'")
	o := flag.String("o", "table", "[Required param] - Output in an specific format. Usage: '-o [ table | yaml | json ]'")
	h := flag.Bool("h", false, "Help usage example: './pullsecret-validator-command -f ./pull-secret.json -o table' ")
	flag.Parse()

	if *h || *f == "" {
		fmt.Println("This is a tool to validate the Pull Secret file from the command line interface.")
		fmt.Println("-- Input Required -- The Pull Secret File as well as the output format ")
		fmt.Println("-- Output  -- The output will be shown in the output format selected with the param -o. By default, the value is table.")
		fmt.Println()
		fmt.Println("Usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	command.ValidatePullSecret(*f, *o)
}
