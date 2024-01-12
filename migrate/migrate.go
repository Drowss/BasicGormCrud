package migrate

import (
	"pruebas/configs"
	"pruebas/models"
)

func SyncDataBase() {
	configs.DB.AutoMigrate(&models.Person{})
}
