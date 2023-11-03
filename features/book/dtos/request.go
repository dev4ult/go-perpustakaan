package dtos

type InputBook struct {
	Title           string `json:"title" form:"title" validate:"required"`
	Summary         string `json:"summary" form:"summary" validate:"required"`
	PublicationYear int    `json:"publication-year" form:"publication-year" validate:"required"`
	Quantity        int    `json:"qty" form:"qty" validate:"required"`
	Language        string `json:"language" form:"language" validate:"required"`
	NumberOfPages   int    `json:"number-of-pages" form:"number-of-pages" validate:"required"`
	CategoryID      int    `json:"category-id" form:"category-id" validate:"required"`
	PublisherID     int    `json:"publisher-id" form:"publisher-id" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}