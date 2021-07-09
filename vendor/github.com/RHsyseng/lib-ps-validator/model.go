package lib_ps_validator

type Payload struct {
	Auths map[string]struct {
		Auth  string `json:"auth"`
		Email string `json:"email,omitempty"`
	}
}

type WebData struct {
	Input     interface{}
	ResultOK  interface{}
	ResultKO  interface{}
	ResultCon interface{}
}
