package dtos


type ResAuthor struct {
	FullName string `json:"full-name"`
	DOB string `json:"dob"`
	Biography string `json:"biography"`
}

type BookAuthors struct {
	Title string `json:"title"`
	CoverImage string `json:"cover-img"`
	Summary string `json:"summary"`
	PublicationYear int `json:"publication-year"`
	Quantity int `json:"qty"`
	Language string `json:"language"`
	NumberOfPages int `json:"number-of-pages"`

	Authors []ResAuthor
}
