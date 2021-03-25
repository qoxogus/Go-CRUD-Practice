package controller

import (
	"fmt"
	"goCRUD/database"

	"github.com/gin-gonic/gin"
)

type InsertParam struct {
	Name      string `json:"name" form:"name" query:"name"`
	ID        string `json:"id" form:"id" query:"id"`
	Pw        string `json:"pw" form:"pw" query:"pw"`
	IsManager bool   `json:"ismanager" form:"ismanager" query:"ismanager"`
}

func Create(c *gin.Context) {
	u := new(InsertParam)

	if err := c.Bind(u); err != nil {
		return
	}

	fmt.Println(u.Name, u.ID, u.Pw)

	User := &database.Crud{ID: u.ID, PW: u.Pw, Name: u.Name, IsManager: u.IsManager}
	err := database.DB.Create(User).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Insert 실패",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":    200,
		"message":   "Insert가 완료되었습니다.",
		"Name":      u.Name,
		"ID":        u.ID,
		"PW":        u.Pw,
		"Ismanager": u.IsManager,
	})
	return
}
