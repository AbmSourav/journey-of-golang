# Render HTML file



``` go
func main() {
	// we got index, and about two html file in /static folder
	// so routes are: / , /about.html
	http.Handle("/", http.FileServer(http.Dir("./webDev/besics/renderStaticHTML/static/")))
	http.ListenAndServe(":3050", nil)
}
```