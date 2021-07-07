package command

import (
	"encoding/json"
	v "github.com/RHsyseng/lib-ps-validator"
	"io/ioutil"
	"log"
	"os"
)

func readFromFile(file string) ([]byte, error) {

	jsonFile, err := os.Open(file)

	if err != nil {
		log.Println(string(ColorRed), "Not possible open the Pull Secret file.")
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	err = validateJsonFormat(byteValue)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}

func validateJsonFormat(input []byte) error {
	var payload v.Payload

	err := json.Unmarshal(input, &payload)

	if err != nil {
		log.Println(string(ColorRed), "Not valid Json Format file...")
		return err
	}
	return nil
}
