package lib_ps_validator

import (
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
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

	req, err := http.NewRequest("GET", "https://"+url+"/v2/auth", nil)
	if err != nil {
		return err, RES_CONERROR
	}

	s := strings.Split(auth, ":")
	req.SetBasicAuth(s[0], s[1])

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 15 * time.Second}
	resp, err2 := client.Do(req)

	if err2 != nil {
		return err2, RES_CONERROR
	} else if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		return nil, RES_VALID
	} else if resp.StatusCode == http.StatusNotFound {
		reqWithoutAuth, _ := http.NewRequest("GET", "https://"+url+"/v2/", nil)
		reqWithoutAuth.SetBasicAuth(s[0], s[1])
		respWithoutAuth, errWithoutAuth := client.Do(reqWithoutAuth)
		if errWithoutAuth != nil || respWithoutAuth.StatusCode == http.StatusNotFound {
			return err2, RES_CONERROR
		} else if errWithoutAuth == nil && (respWithoutAuth.StatusCode == http.StatusOK || respWithoutAuth.StatusCode == http.StatusAccepted) {
			return nil, RES_VALID
		} else if errWithoutAuth == nil && (respWithoutAuth.StatusCode == http.StatusUnauthorized || respWithoutAuth.StatusCode == http.StatusForbidden) {
			return nil, RES_EXPIRED
		}
		return nil, RES_CONERROR
	} else if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return nil, RES_EXPIRED
	} else {
		return err, RES_CONERROR
	}
}
