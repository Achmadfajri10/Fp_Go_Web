package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"Fp_Go_Web/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Cek file .env, jika tidak maka buat file .env
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Missing .env file, Creating one...\n")
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter MySQL username(default: root): ")
		dbUser, _ := reader.ReadString('\n')
		dbUser = strings.TrimSpace(dbUser)
		if dbUser == "" {
			dbUser = "root"
		}

		fmt.Print("Enter MySQL password: ")
		dbPass, _ := reader.ReadString('\n')
		dbPass = strings.TrimSpace(dbPass)

		fmt.Print("Enter MySQL host (default: 127.0.0.1): ")
		dbHost, _ := reader.ReadString('\n')
		dbHost = strings.TrimSpace(dbHost)
		if dbHost == "" {
			dbHost = "127.0.0.1"
		}

		fmt.Print("Enter MySQL port (default: 3306): ")
		dbPort, _ := reader.ReadString('\n')
		dbPort = strings.TrimSpace(dbPort)
		if dbPort == "" {
			dbPort = "3306"
		}

		fmt.Print("Enter database name (default: go_products): ")
		dbName, _ := reader.ReadString('\n')
		dbName = strings.TrimSpace(dbName)
		if dbName == "" {
			dbName = "go_products"
		}

		envContent := fmt.Sprintf("DB_USER=%s\nDB_PASSWORD=%s\nDB_HOST=%s\nDB_PORT=%s\nDB_NAME=%s\nSECRET=auth-api-jwt-secret", dbUser, dbPass, dbHost, dbPort, dbName)
		if err := os.WriteFile(".env", []byte(envContent), 0644); err != nil {
			fmt.Println("Error creating .env file:", err)
			return
		}

		fmt.Println(".env file created successfully.")

		if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env:", err)
			return
		}
	}

	// konek ke database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		fmt.Println("Trying to solve the error by creating the database...")

		// Jika database tidak ditemukan, maka buat database baru
		dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

		tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
		if err != nil {
			fmt.Println("Error creating temporary database connection:", err)
			return
		}

		sqlDB, err := tempDB.DB()
		if err != nil {
			fmt.Println("Error getting SQL DB:", err)
			return
		}

		defer sqlDB.Close()

		_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", os.Getenv("DB_NAME")))
		if err != nil {
			fmt.Println("Error creating database:", err)
			return
		}
		tempDB.Exec("USE " + os.Getenv("DB_NAME"))

		// Coba konek lagi ke database
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Error connecting to database after creation:", err)
			return
		}
	}

	DB = db

	// menjalankan migrasi menggunakan GORM AutoMigrate
	err = db.AutoMigrate(&entities.Category{}, &entities.Product{}, &entities.User{})
	if err != nil {
		fmt.Println("Error migrating schema:", err)
		return
	}

	fmt.Println("Schema migrated successfully!")

}
