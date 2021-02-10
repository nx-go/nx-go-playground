package user

import (
	"api/prisma/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func User(name string) string {
	result := "User " + name
	return result
}

func UsersHandler(ctx context.Context, client db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post, err := client.User.FindMany().Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}
		result, _ := json.MarshalIndent(post, "", "  ")
		fmt.Fprintf(w, "%s\n", result)
	}
}
