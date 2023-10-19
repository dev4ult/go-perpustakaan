package dtos

type ResAuthorization struct {
	FullName string `json:"full-name"`
	AccessToken string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

type User struct {
	ID int
	FullName string 
	Password string 
}
