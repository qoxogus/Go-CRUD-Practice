package controller

import (
	"goCRUD/database"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	Crud := &database.Crud{} //테이블

	err := database.DB.Delete(&Crud, "7").Error //idx == 7인 컬럼 지우기
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "500",
			"message": "delete 실패",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "delete 성공",
	})
	return
}
