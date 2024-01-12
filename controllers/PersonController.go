package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pruebas/configs"
	"pruebas/models"
)

type PersonRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   uint   `json:"phone"`
}

func CreatePerson(c *gin.Context) {
	body := PersonRequestBody{}
	c.BindJSON(&body)
	person := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}
	result := configs.DB.Create(&person)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, &person)
}

func GetPersons(c *gin.Context) {
	var persons []models.Person
	configs.DB.Find(&persons)
	c.JSON(http.StatusOK, &persons)
	return
}

func GetPersonById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)
	c.JSON(http.StatusOK, &person)
	return
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)
	body := PersonRequestBody{}
	c.BindJSON(&body)
	data := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}
	configs.DB.Model(&person).Updates(data)
	c.JSON(http.StatusOK, &person)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.Delete(&person, id)
	c.JSON(http.StatusOK, gin.H{
		"deleted": true,
	})
	return
}
