package dtos

type InputAuthor struct {
	FullName string `json:"full-name" form:"full-name" validate:"required"`
	DOB string `json:"dob" form:"dob" validate:"required"`
	Biography string `json:"biography" form:"biography" validate:"required"`
}

type InputAuthorshipIDS struct {
	AuthorID int `json:"author-id" form:"author-id" validate:"required"`
	BookID int `json:"book-id" form:"book-id" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}