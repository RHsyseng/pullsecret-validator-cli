package command

import (
	v "github.com/RHsyseng/lib-ps-validator"
	"log"
	"os"
)
const(
	OUTPUT_TABLE 	= "table"
	OUTPUT_JSON		= "json"
	OUTPUT_YAML		= "yaml"
)

func ValidatePullSecret(file string, outputFormat string) {

	input, err := readFromFile(file)
	if err != nil {
		log.Fatal(string(ColorRed), "Not possible validate the Pull Secret due to: "+err.Error())
		os.Exit(500)
	}
	log.Println("Starting the Pull Secret file validation...It could take a time!")
	result := v.Validate(input)

	switch outputFormat {
	case OUTPUT_TABLE:
		writeOutputTable(result)
	case OUTPUT_JSON:
		writeOutputJson(result)
	case OUTPUT_YAML:
		writeOutputYaml(result)
	default:
		writeOutputTable(result)
	}


}
