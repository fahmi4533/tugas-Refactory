package productcontroller

import (
	"encoding/json"
	"fahmi/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var produk []models.Produk

	models.DB.Find(&produk)
	c.JSON(http.StatusOK, gin.H{"produk": produk})
}
func Show(c *gin.Context) {
	var produk models.Produk

	id := c.Param("id")

	if err := models.DB.First(&produk, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tak ada woy!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"produk w": produk})
}
func Create(c *gin.Context) {
	var produk models.Produk

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}
	models.DB.Create(&produk)
	c.JSON(http.StatusOK, gin.H{"peroduk": produk})
}
func Update(c *gin.Context) {
	var produk models.Produk

	id := c.Param("id")

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&produk).Where("id = ?", id).Updates(&produk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengubah data bro!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "berhasil bro"})
}
func Delete(c *gin.Context) {
	var produk models.Produk

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()

	if models.DB.Delete(&produk, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Apa yang kamu hapus!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "mantap bro"})
}
