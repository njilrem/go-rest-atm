package dto

//Login credential
type AuthCredentials struct {
	cardNum string `json:"cardNum"`
	exprireDate string `json:"expireDate"`
	cvv2 string `json:"cvv2"`
}