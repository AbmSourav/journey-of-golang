# Struct

A struct is a collection of fields. 
``` go
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
```
<br>

Struct fields also can be accessed through a struct pointer. 
``` go
article := &post
fmt.Println((*article).title)
```
