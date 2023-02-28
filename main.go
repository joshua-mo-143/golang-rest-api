package main

import	("fmt"
	"log"
	 "github.com/gofiber/fiber"
	 "github.com/joho/godotenv"
	)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}

	app := fiber.New()
	
	app.Get("/", hello_world)
	app.Get("/meme", meme)

	fmt.Println("Server is running on port 5000...")
	
	app.Listen(5000)
}
