package main

import (
	"flag"
	"fmt"
	"github.com/RHsyseng/pullsecret-validator-cli/command"
	"os"
)

func main() {

	f := flag.String("f", "", "[Required param] - Pull Secret file in json format. Usage: '-f ./pull-secret.json'")
	flag.Bool("r", false, "[Optional param] - Array of registry entries from the pull secret file to be validated. Usage: '-r quay.io cloud.openshift.io'")
	h := flag.Bool("h", false, "Help usage example: './pullsecret-validator-command -f ./pull-secret.json -r quay.io cloud.openfhit.io' ")
	flag.Parse()
	tail := flag.Args()

	if *h || *f == "" {
		fmt.Println("This is a tool to validate the Pull Secret file from the command line interface.")
		fmt.Println("-- Input Required -- The Pull Secret File ")
		fmt.Println("-- Input Optional -- An array of entries to be only validated as a subset of the whole Pull Secret file.")
		fmt.Println("-- Output  -- The output will be true when all entries will be validated. If some entries have expired will be shown as an output table")
		fmt.Println()
		fmt.Println("Usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	command.ValidatePullSecret(*f, tail)
}
