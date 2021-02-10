package feature

import (
	data "api/libs/api/post/data-access"
	"api/prisma/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func PostFeature(name string) string {
	result := "PostFeature " + name
	return result
}

func PostsHandler(ctx context.Context, client db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts := data.GetPosts(ctx, client)
		result, _ := json.MarshalIndent(posts, "", "  ")
		fmt.Fprintf(w, "%s\n", result)
	}
}

func PostHandler(ctx context.Context, client db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		post := data.GetPost(ctx, client, vars["postId"])

		if post == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Post not found")
			return
		}
		result, _ := json.MarshalIndent(post, "", "  ")
		fmt.Fprintf(w, "%s\n", result)
	}
}
