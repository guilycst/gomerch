package helpers

import (
	"gomerch/db"
	"gomerch/models"
)

func Populate() {
	gdb := db.GetConnection()
	// Migrate the schema
	gdb.AutoMigrate(&models.Product{})

	//Clear
	var p []models.Product
	gdb.Where("1=1").Delete(&p)

	// Create
	gdb.Create(&models.Product{Name: "Shirt", Description: "Black", Price: 19, Amount: 99})
	gdb.Create(&models.Product{Name: "Headset", Description: "Gamer", Price: 99, Amount: 10})
	gdb.Create(&models.Product{Name: "Mouse", Description: "LED", Price: 30, Amount: 100})
	gdb.Create(&models.Product{Name: "Keyboard", Description: "Mechanical", Price: 50, Amount: 845})
}
