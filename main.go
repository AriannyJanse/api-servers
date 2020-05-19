package main

import (
	"api-server/models"
	"api-server/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func main() {
	s := server.New()
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	models.Init()
	fmt.Println("Connection started on port:", port)
	log.Fatal(fasthttp.ListenAndServe(":"+port, s.Router().Handler))
}
