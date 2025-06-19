package smsactivate

type NumberData struct {
	ActivationID       string  `json:"activationId"`
	PhoneNumber        string  `json:"phoneNumber"`
	ActivationCost     float64 `json:"activationCost"`
	Currency           int     `json:"currency"`
	CountryCode        string  `json:"countryCode"`
	CanGetAnotherSms   bool    `json:"canGetAnotherSms"`
	ActivationTime     string  `json:"activationTime"`
	ActivationEndTime  string  `json:"activationEndTime"`
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
