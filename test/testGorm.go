package main

import (
	"fmt"
	"gin_chat/models"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Name = "Aaron"

	// Create
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // find product with integer primary key
	//db.First(user, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(user).Update("Password", "123")
	// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&product, 1)
}
