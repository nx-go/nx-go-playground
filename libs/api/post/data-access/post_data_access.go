package post_data_access

import (
	"api/prisma/db"
	"context"
	"log"
)

func PostDataAccess(name string) string {
	result := "PostDataAccess " + name
	return result
}

func GetPosts(ctx context.Context, client db.PrismaClient) []db.PostModel {
	result, err := client.Post.FindMany().Exec(ctx)
	if err != nil {
		log.Printf("Prisma error: %s\n", err)
	}
	return result
}

func GetPost(ctx context.Context, client db.PrismaClient, postId string) *db.PostModel {
	result, err := client.Post.FindUnique(db.Post.ID.Equals(postId)).Exec(ctx)
	if err != nil {
		log.Printf("Prisma error: %s\n", err)
	}
	return result
}
