package smsactivate

type NumberData struct {
	Status             string `json:"status"`
	ActivationID       string `json:"activationId"`
	PhoneNumber        string `json:"phoneNumber"`
	ActivationCost     string `json:"activationCost"`
	CountryCode        string `json:"countryCode"`
	CanGetAnotherSms   bool   `json:"canGetAnotherSms"`
	ActivationTime     string `json:"activationTime"`
	ActivationEndTime  string `json:"activationEndTime"`
	ActivationOperator string `json:"activationOperator"`
	*Error
}

type Error struct {
	Msg      string `json:"msg"`
	ErrorMsg string `json:"errorMsg"`
	Info     struct {
		Min float64 `json:"min"`
	} `json:"info"`
}
