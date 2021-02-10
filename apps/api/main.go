package main

import (
	"api/libs/api/core"
	post "api/libs/api/post/feature"
	"api/libs/api/user"
	"api/prisma/db"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Hello returns hello with the name you pass in :)
func Hello(name string) string {
	result := "Hello " + name
	fmt.Println(result)
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	var input = "World"
	if len(r.URL.Path) > 1 {
		input = r.URL.Path[1:]
	}
	fmt.Fprintf(w, "%s\n", Hello(input))
}

func getAppRouter(ctx context.Context, client *db.PrismaClient) http.Handler {

	appRouter := mux.NewRouter().StrictSlash(true)

	appRouter.HandleFunc("/", handler)
	appRouter.HandleFunc("/uptime", core.UptimeHandler)
	appRouter.HandleFunc("/users", user.UsersHandler(ctx, *client))
	appRouter.HandleFunc("/posts", post.PostsHandler(ctx, *client))
	appRouter.HandleFunc("/posts/{postId}", post.PostHandler(ctx, *client))
	return appRouter
}

func main() {
	var defaultHost = "localhost"
	host := os.Getenv("HOST")
	if host == "" {
		host = defaultHost
	}

	var defaultPort = "3000"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	address := fmt.Sprintf("%s:%s", host, port)

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	fmt.Printf("Listening on port %s\n", address)
	log.Fatal(http.ListenAndServe(address, getAppRouter(ctx, client)))
}
