package lib_ps_validator

//Payload Struct to represent the auth pull secret file structure
type Payload struct {
	Auths map[string]struct {
		Auth  string `json:"auth"`
		Email string `json:"email,omitempty"`
	}
}

//WebData is the representation of the results obtained into the validation process
type WebData struct {
	Input     interface{}
	ResultOK  interface{}
	ResultKO  interface{}
	ResultCon interface{}
}
