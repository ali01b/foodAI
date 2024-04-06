package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Projede .env dosyası bulunamadı!")
	}
}
