package smsactivate

type NumberData struct {
	ActivationID       int    `json:"activationId"`
	PhoneNumber        string `json:"phoneNumber"`
	ActivationCost     string `json:"activationCost"`
	CountryCode        string `json:"countryCode"`
	CanGetAnotherSms   string `json:"canGetAnotherSms"`
	ActivationTime     string `json:"activationTime"`
	ActivationOperator string `json:"activationOperator"`
}
