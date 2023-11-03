package dtos

type ResMember struct {
	FullName string `json:"full_name"`
	CredentialNumber string `json:"credential_number"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
}
