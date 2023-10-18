package dtos

type ResLogin struct {
	FullName string `json:"full-name"`
	AccessToken string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}
