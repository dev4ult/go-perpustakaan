package dtos

type ResMember struct {
	FullName string `json:"full-name"`
	CredentialNumber string `json:"credential-number"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone-number"`
	Address string `json:"address"`
}
