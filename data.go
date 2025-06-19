package smsactivate

type NumberData struct {
	ActivationID       int     `json:"activationId"`
	PhoneNumber        string  `json:"phoneNumber"`
	ActivationCost     float64 `json:"activationCost"`
	Currency           int     `json:"currency"`
	CountryCode        string  `json:"countryCode"`
	CanGetAnotherSms   string  `json:"canGetAnotherSms"`
	ActivationTime     string  `json:"activationTime"`
	ActivationOperator string  `json:"activationOperator"`
	*Error
}

type Error struct {
	Msg      string `json:"msg"`
	ErrorMsg string `json:"errorMsg"`
	Info     struct {
		Min float64 `json:"min"`
	} `json:"info"`
}
