package main

import (
	"log"
	"os"

	blogposts "github.com/zshainsky/learning_go_with_testing/filesystem/v1/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
