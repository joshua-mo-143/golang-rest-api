package main

import	("fmt"
	"context"
	"os"
	 "github.com/gofiber/fiber"
	"github.com/jackc/pgx/v5/pgxpool"
	)

func hello_world(c *fiber.Ctx) {
	c.Send("Hello world")
}

func meme(c *fiber.Ctx) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	var message string
	err = conn.QueryRow(context.Background(), "select message from messages").Scan(&message)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	c.Send(message)
	
}
