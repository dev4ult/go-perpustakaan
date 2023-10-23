package dtos

type ResAuthorization struct {
	FullName string `json:"full-name"`
	AccessToken string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

type ResLibrarian struct {
	FullName string `json:"full-name"`
	StaffID string `json:"staff-id"`
	NIK int `json:"nik"`
	PhoneNumber string `json:"phone-number"`
	Address string `json:"address"`
	Email string `json:"email"`
}

type Token struct {
	AccessToken string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

type User struct {
	ID int
	FullName string 
	Password string 
}
