package db

import (
	"fmt"
	"log"
	"os"
	"webscaper/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(envPath string) {

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connectionstring := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	database, err := gorm.Open(postgres.Open(connectionstring), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database

	DB.AutoMigrate(&models.Monitor{})

}

func getRecords(url string) ([]models.Monitor, error) {

	var records []models.Monitor

	err := DB.Where("url = ?", url).Order("created_at_desc").Find(&records).Error

	if err != nil {
		return nil, err
	}

	return records, nil

}

func calculateChangePercent(oldText, newText string) float64 {

	if oldText == "" {
		return 100
	}

	if oldText == newText {
		return 0
	}

	diff := 0
	minLen := len(oldText)

	if len(newText) < minLen {
		minLen = len(newText)
	}

	for i := 0; i < minLen; i++ {
		if oldText[i] != newText[i] {
			diff++
		}
	}

	return (float64(diff) / float64(len(oldText))) * 100
}

func BuildComparison(url string, currentTitle string, currentBody string) ([]models.MonitorView, error) {

	records, err := getRecords(url)
	if err != nil {
		return nil, err
	}

	var result []models.MonitorView

	for _, r := range records {

		titleChange := calculateChangePercent(r.HashValuesTitle, currentTitle)
		bodyChange := calculateChangePercent(r.HashValuesBody, currentBody)

		view := models.MonitorView{
			Time:        r.CreatedAt.Format("2006-01-02 15:04"),
			TitleChange: titleChange,
			BodyChange:  bodyChange,
		}

		result = append(result, view)
	}

	return result, nil
}
