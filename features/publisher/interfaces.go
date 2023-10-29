package publisher

import (
	"perpustakaan/features/publisher/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page int, size int, searchKey string) ([]Publisher, error)
	Insert(newPublisher Publisher) (int, error)
	SelectByID(publisherID int) (*Publisher, error)
	Update(publisher Publisher) (int, error)
	DeleteByID(publisherID int) (int, error)
}

type Usecase interface {
	FindAll(page int, size int, searchKey string) ([]dtos.ResPublisher, string)
	FindByID(publisherID int) (*dtos.ResPublisher, string)
	Create(newPublisher dtos.InputPublisher) (*dtos.ResPublisher, string)
	Modify(publisherData dtos.InputPublisher, publisherID int) (bool, string)
	Remove(publisherID int) (bool, string)
}

type Handler interface {
	GetPublishers() echo.HandlerFunc
	PublisherDetails() echo.HandlerFunc
	CreatePublisher() echo.HandlerFunc
	UpdatePublisher() echo.HandlerFunc
	DeletePublisher() echo.HandlerFunc
}
