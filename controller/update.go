package controller

import (
	"fmt"
	"goCRUD/database"

	"github.com/gin-gonic/gin"
)

type UpdateParam struct {
	ID   string `json:"id" form:"id" query:"id"`
	Pw   string `json:"pw" form:"pw" query:"pw"`
	Repw string `json:"repw" form:"repw" query:"repw"`
}

func Update(c *gin.Context) {
	u := new(UpdateParam)

	if err := c.Bind(u); err != nil {
		return
	}

	fmt.Println(u.ID, u.Pw, u.Repw)

	Crud := &database.Crud{} //테이블

	err := database.DB.Where("id = ? AND pw = ?", u.ID, u.Pw).Find(Crud).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "일치하지 않습니다.",
		})
		return
	}

	fmt.Println(Crud)

	err = database.DB.Model(&Crud).Update("pw", u.Repw).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "업데이트 실패.",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "업데이트 성공.",
	})
	return
}
