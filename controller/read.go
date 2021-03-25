package controller

import (
	"fmt"
	"goCRUD/database"

	"github.com/gin-gonic/gin"
)

type SelectParam struct {
	ID string `json:"id" form:"id" query:"id"`
	Pw string `json:"pw" form:"pw" query:"pw"`
}

func Read(c *gin.Context) {
	u := new(SelectParam)

	if err := c.Bind(u); err != nil {
		return
	}

	Crud := &database.Crud{} //테이블

	database.DB.Find(&Crud) //select

	fmt.Println(Crud) //selct value

	err := database.DB.Where("id = ? AND pw = ?", u.ID, u.Pw).Find(Crud).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "일치하지 않습니다.",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":       500,
		"message":      "일치합니다.",
		"DB idx":       Crud.Idx,
		"DB Name":      Crud.Name,
		"DB id":        Crud.ID,
		"DB pw":        Crud.PW,
		"DB ismanager": Crud.IsManager,
	})
}
