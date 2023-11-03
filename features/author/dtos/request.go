package dtos

type InputAuthor struct {
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	DOB string `json:"dob" form:"dob" validate:"required"`
	Biography string `json:"biography" form:"biography" validate:"required"`
}

type InputAuthorshipIDS struct {
	AuthorID int `json:"author_id" form:"author_id" validate:"required"`
	BookID int `json:"book_id" form:"book_id" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}