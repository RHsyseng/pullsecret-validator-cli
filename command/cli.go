package command

import (
	v "github.com/RHsyseng/lib-ps-validator"
	"log"
	"os"
)

func ValidatePullSecret(file string, registries []string) {

	input, err := readFromFile(file)
	if err != nil {
		log.Fatal(string(ColorRed), "Not possible validate the Pull Secret due to: "+err.Error())
		os.Exit(500)
	}
	log.Println("Starting the Pull Secret file validation...It could take a time!")
	result := v.Validate(input)

	writeOutputTable(result, registries)

}
