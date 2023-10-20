package dtos

type ResBook struct {
	Title           string `json:"title"`
	CoverImage      string `json:"cover-img"`
	Summary         string `json:"summary"`
	PublicationYear int    `json:"publication-year"`
	Quantity        int    `json:"qty"`
	Language        string `json:"language"`
	NumberOfPages   int    `json:"number-of-pages"`

	Category  string `json:"category"`
	Publisher string `json:"publisher"`
}

type AfterInsert struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	CoverImage      string `json:"cover-img"`
	Summary         string `json:"summary"`
	PublicationYear int    `json:"publication-year"`
	Quantity        int    `json:"qty"`
	Language        string `json:"language"`
	NumberOfPages   int    `json:"number-of-pages"`

	CategoryID  string `json:"category-id"`
	PublisherID string `json:"publisher-id"`
}