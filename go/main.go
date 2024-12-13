package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Album struct {
// 	gorm.Model
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Printf("main---->\n")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
	fmt.Printf("<----main\n")
}

// // Migrate the schema
// func Migrate() {
//         db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//         if err != nil {
//                 panic("failed to connect database")
//         }
//
//         db.AutoMigrate(&Album{})
// }
//
// // Create
// func Create(album Album) {
//         db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//         if err != nil {
//                 panic("failed to connect database")
//         }
//
//         db.Create(&album)
// }
//
// // Read
// func Read() {
//         db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//         if err != nil {
//                 panic("failed to connect database")
//         }
//
//         var album Album
//         db.First(&album, 1)
// }
//
// // Update
// func UpdateById(id string, album Album) {
//         db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//         if err != nil {
//                 panic("failed to connect database")
//         }
//
//         var oldAlbum Album
//         oldAlbum.ID = id
//         db.Model(oldAlbum).Updates(album)
// }
//
// // Delete
// func Delete(album Album) {
//         db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//         if err != nil {
//                 panic("failed to connect database")
//         }
//
//         db.Delete(&album, 1)
// }

// type Product struct {
//      gorm.Model
//      Code  string
//      Price uint
// }
//
// func test() {
//      db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//      if err != nil {
//              panic("failed to connect database")
//      }
//
//      // Migrate the schema
//      db.AutoMigrate(&Product{})
//
//      // Create
//      db.Create(&Product{Code: "D42", Price: 100})
//
//      // Read
//      var product Product
//      db.First(&product, 1)
//      db.First(&product, "code = ?", "D42")
//
//      // Update
//      db.Model(&product).Update("Price", 200)
//      // Update
//      db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
//      db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
//
//      // Delete
//      db.Delete(&product, 1)
// }
