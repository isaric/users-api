package util

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"users-api/models"
)

func GetId(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}

func Encode(w http.ResponseWriter, x interface{}) {
	json.NewEncoder(w).Encode(x)
}

var Database *gorm.DB

func OpenConnectionAndSetupSchema() {
	var url string = os.Getenv("DB_HOST")
	if len(url) < 1 {url = "localhost" }
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=5432 user=postgres dbname=postgres sslmode=disable password=postgres", url)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable(&models.User{}, &models.Group{})

	err = db.AutoMigrate(&models.User{}, &models.Group{})
	if err != nil {
		panic("failed to migrate database")
	}

	Database = db
}

func HandleJsonError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "bad payload", http.StatusBadRequest)
	}
}
