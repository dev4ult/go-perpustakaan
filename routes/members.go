package routes

import (
	"perpustakaan/features/member"
	"perpustakaan/features/member/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Members(e *echo.Echo, handler member.Handler) {
	members := e.Group("/members")

	members.GET("", handler.GetMembers(), m.Authorization("librarian"))
	members.POST("", handler.CreateMember(), m.RequestValidation(dtos.InputMember{}))
	
	members.GET("/:id", handler.MemberDetails(), m.Authorization("librarian"))
	members.PUT("/:id", handler.UpdateMember(), m.Authorization("librarian"), m.RequestValidation(dtos.InputMember{}))
	members.DELETE("/:id", handler.DeleteMember(), m.Authorization("librarian"))
}