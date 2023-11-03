package dtos


type ResAuthor struct {
	FullName string `json:"full_name"`
	DOB string `json:"dob"`
	Biography string `json:"biography"`
}

type BookAuthors struct {
	Title string `json:"title"`
	CoverImage string `json:"cover_img"`
	Summary string `json:"summary"`
	PublicationYear int `json:"publication_year"`
	Quantity int `json:"qty"`
	Language string `json:"language"`
	NumberOfPages int `json:"number_of_pages"`

	Authors []ResAuthor
}
