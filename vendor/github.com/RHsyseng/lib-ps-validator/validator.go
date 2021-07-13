package lib_ps_validator

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/containers/image/v5/docker"
	"strings"
)

const (
	RES_VALID    = "valid"
	RES_EXPIRED  = "expired"
	RES_CONERROR = "conn-error"
	ERROR        = "error"
)

func Validate(input []byte) WebData {

	var payload Payload

	err := json.Unmarshal(input, &payload)
	if err != nil {
		fmt.Println(ERROR, err)
	}

	var resultKOConArray string
	var resultKOArray string
	var resultOKArray string

	for k, v := range payload.Auths {

		sDec, err := b64.StdEncoding.DecodeString(v.Auth)
		if err != nil {
			return WebData{input, "", "", ""}
		}
		auth := string(sDec)

		err, res := loginToRegistry(k, auth)
		if err != nil || res == RES_CONERROR {
			resultKOConArray += k + "\n"
		} else if res == RES_VALID {
			resultOKArray += k + "\n"
		} else if res == RES_EXPIRED {
			resultKOArray += k + "\n"
		}
	}
	return WebData{input, resultOKArray, resultKOArray, resultKOConArray}
}

func loginToRegistry(url, auth string) (error, string) {
	if url != "" {
		s := strings.Split(auth, ":")
		ctx := context.Background()
		err := docker.CheckAuth(ctx, nil, s[0], s[1], url)
		if err != nil {
			if strings.Contains(err.Error(), "no such host") {
				return err, RES_CONERROR
			}
			return nil, RES_EXPIRED
		}
		return nil, RES_VALID
	}
	return errors.New("URL is empty"), RES_CONERROR
}
