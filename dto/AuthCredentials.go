package dto

//Login credential
type AuthCredentials struct {
	CardNum     string `json:"cardNum"`
	ExpireDate  string `json:"expireDate"`
	Cvv2        string `json:"cvv2"`
}

//Login credential
type AuthCredentialsPIN struct {
	CardNum     string `json:"cardNum"`
	Pin         string `json:"pin"`
}

type AuthAdminCredentials struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type AuthTransaction struct {
	ID string `json:"card_id"`
}