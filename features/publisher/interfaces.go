package publisher

import (
	"perpustakaan/features/publisher/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Publisher
	Insert(newPublisher Publisher) int64
	SelectByID(publisherID int) *Publisher
	Update(publisher Publisher) int64
	DeleteByID(publisherID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResPublisher
	FindByID(publisherID int) *dtos.ResPublisher
	Create(newPublisher dtos.InputPublisher) *dtos.ResPublisher
	Modify(publisherData dtos.InputPublisher, publisherID int) bool
	Remove(publisherID int) bool
}

type Handler interface {
	GetPublishers() echo.HandlerFunc
	PublisherDetails() echo.HandlerFunc
	CreatePublisher() echo.HandlerFunc
	UpdatePublisher() echo.HandlerFunc
	DeletePublisher() echo.HandlerFunc
}
