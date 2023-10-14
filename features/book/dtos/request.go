package dtos

type InputBook struct {
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}