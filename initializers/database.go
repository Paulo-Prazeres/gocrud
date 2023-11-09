package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Code  string
	Price uint
}

var DB *gorm.DB

func ConnectToDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	} else {
		println("Connected to database")
	}
	// db.AutoMigrate(&Product{})

	// insertProduct := &Product{Code: "D42", Price: 100}

	// db.Create(insertProduct)
	// fmt.Printf("insert ID: %d, Code: %s, Price: %d\n",
	// 	insertProduct.ID, insertProduct.Code, insertProduct.Price)

	// readProduct := &Product{}
	// db.First(&readProduct, "code = ?", "D42") // find product with code D42

	// fmt.Printf("read ID: %d, Code: %s, Price: %d\n",
	// 	readProduct.ID, readProduct.Code, readProduct.Price)
}
