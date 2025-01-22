package controllers

import (
	"net/http"

	"github.com/enjiar/crud-gps/config"
	"github.com/enjiar/crud-gps/models"
	"github.com/gin-gonic/gin"
)

func GetAllDevices(c *gin.Context) {
	var devices []models.GPSDevice
	config.DB.Find(&devices)
	c.JSON(http.StatusOK, devices)
}

func CreateDevice(c *gin.Context) {
	var device models.GPSDevice
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&device)
	c.JSON(http.StatusCreated, device)
}

func GetDeviceByID(c *gin.Context) {
	var device models.GPSDevice
	id := c.Param("id")

	if err := config.DB.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	c.JSON(http.StatusOK, device)
}

func UpdateDevice(c *gin.Context) {
	var device models.GPSDevice
	id := c.Param("id")

	if err := config.DB.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&device)
	c.JSON(http.StatusOK, device)
}

func DeleteDevice(c *gin.Context) {
	var device models.GPSDevice
	id := c.Param("id")

	if err := config.DB.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	config.DB.Delete(&device)
	c.JSON(http.StatusOK, gin.H{"message": "Device deleted successfully"})
}
