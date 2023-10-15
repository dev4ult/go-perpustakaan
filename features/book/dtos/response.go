package dtos

type ResBook struct {
	Title           string `json:"title"`
	Summary         string `json:"summary"`
	PublicationYear string `json:"publication-year"`
	Quantity        int    `json:"qty"`
	Language        string `json:"language"`
	NumberOfPages   int    `json:"number-of-pages"`

	CategoryID  int `json:"category-id"`
	PublisherID int `json:"publisher-id"`
	FineTypeID  int `json:"fine-type-id"`
}