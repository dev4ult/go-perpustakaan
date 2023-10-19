package routes

import (
	"perpustakaan/features/member"

	"github.com/labstack/echo/v4"
)

func Members(e *echo.Echo, handler member.Handler) {
	members := e.Group("/members")

	members.GET("", handler.GetMembers())
	members.POST("", handler.CreateMember())
	
	members.GET("/:id", handler.MemberDetails())
	members.PUT("/:id", handler.UpdateMember())
	members.DELETE("/:id", handler.DeleteMember())
}