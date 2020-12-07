package main

import "fmt"

type blogPost struct {
	author  string
	title   string
	postId  int  
}
  
func main() {
	post := blogPost{
		author: "Sourav",
		title: "Hello Universe",
		postId: 12345,
	}
	fmt.Printf("Title: %s.\nAuthor: %s", post.title, post.author)
}
