package main

import "github.com/RHsyseng/pullsecret-validator-cli/cli"

func main() {

	f := flag.String("f", "", "Pull secret file to use in json format: -f ./pull-secret.json ")
	h := flag.Bool("h", false, "Help usage example: ./ocp-release -v nightly -c metal-ipi aws gcp")
	flag.Parse()
	tail := flag.Args()

	if *h {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if len(tail) < 1 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	cli.ValidatePullSecret() //validatePullSecret(*f, tail)
}
