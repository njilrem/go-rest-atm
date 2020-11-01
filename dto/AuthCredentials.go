package dto

//Login credential
type AuthCredentials struct {
	CardNum     string `json:"cardNum"`
	ExpireDate  string `json:"expireDate"`
	Cvv2        string `json:"cvv2"`
}

type AuthAdminCredentials struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}