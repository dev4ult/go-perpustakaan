package ayam

import (
	"perpustakaan/features/ayam/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Ayam
	Insert(newAyam Ayam) int64
	SelectByID(ayamID int) *Ayam
	Update(ayam Ayam) int64
	DeleteByID(ayamID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResAyam
	FindByID(ayamID int) *dtos.ResAyam
	Create(newAyam dtos.InputAyam) *dtos.ResAyam
	Modify(ayamData dtos.InputAyam, ayamID int) bool
	Remove(ayamID int) bool
}

type Handler interface {
	GetAyams() echo.HandlerFunc
	AyamDetails() echo.HandlerFunc
	CreateAyam() echo.HandlerFunc
	UpdateAyam() echo.HandlerFunc
	DeleteAyam() echo.HandlerFunc
}
