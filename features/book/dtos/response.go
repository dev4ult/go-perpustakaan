package dtos

type ResBook struct {
	Title           string `json:"title"`
	CoverImage      string `json:"cover_img"`
	Summary         string `json:"summary"`
	PublicationYear int    `json:"publication_year"`
	Quantity        int    `json:"qty"`
	Language        string `json:"language"`
	NumberOfPages   int    `json:"number_of_pages"`

	Category  string `json:"category"`
	Publisher string `json:"publisher"`
}

type AfterInsert struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	CoverImage      string `json:"cover_img"`
	Summary         string `json:"summary"`
	PublicationYear int    `json:"publication_year"`
	Quantity        int    `json:"qty"`
	Language        string `json:"language"`
	NumberOfPages   int    `json:"number_of_pages"`

	CategoryID  int `json:"category_id"`
	PublisherID int `json:"publisher_id"`
}