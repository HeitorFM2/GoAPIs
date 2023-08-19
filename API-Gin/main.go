package main

import (
	"api-gin2/api/routes"
	"api-gin2/configs"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loadgin .env file")
	}
	configs.Connection()
}

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run()

}

// 	configs.DB.AutoMigrate(&models.Users{}) CRIA UMA TABELA NO BANCO ATRAVES DE UMA STRUCT
// (*) busca o valor da memoria alocada
// (&) ponteiro que mostra a memoria que a variavel esta alocada
