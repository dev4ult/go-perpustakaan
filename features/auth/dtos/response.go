package dtos

type ResAuthorization struct {
	FullName string `json:"full_name"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResLibrarian struct {
	FullName string `json:"full_name"`
	StaffID string `json:"staff_id"`
	NIK int `json:"nik"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
	Email string `json:"email"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	ID int
	FullName string 
	Password string 
}
