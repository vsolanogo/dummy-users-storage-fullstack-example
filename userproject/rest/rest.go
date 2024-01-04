package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()

	h, _ := NewHandler()

	userGroup := r.Group("/users")
	{
		userGroup.GET("", h.GetUsers)
		userGroup.POST("", h.AddUser)
		userGroup.GET("/:id", h.GetUser)
		userGroup.DELETE("/:id", h.RemoveUser)
		userGroup.PATCH("/:id", h.PatchUser)
	}

	return r.Run(address)
}
