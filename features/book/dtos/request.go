package dtos

type InputBook struct {
	Title           string `json:"title" form:"title" validate:"required"`
	Summary         string `json:"summary" form:"summary" validate:"required"`
	PublicationYear int    `json:"publication_year" form:"publication_year" validate:"required"`
	Quantity        int    `json:"qty" form:"qty" validate:"required"`
	Language        string `json:"language" form:"language" validate:"required"`
	NumberOfPages   int    `json:"number_of_pages" form:"number_of_pages" validate:"required"`
	CategoryID      int    `json:"category_id" form:"category_id" validate:"required"`
	PublisherID     int    `json:"publisher_id" form:"publisher_id" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}