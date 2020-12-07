# HTML Template

in `Execute` method the second parameter is data. These data is also available in the `html` file. It works like *Mustache* template

``` go
func handeler(resw http.ResponseWriter, req *http.Request) {
	// template := template.Must(template.ParseFiles("./webDev/besics/02.template/public/index.gohtml"))
	template := template.Must(template.ParseFiles("./webDev/besics/02.template/public/index.html"))
	template.Execute(resw, "Universe")
}

func main() {
	http.HandleFunc("/", handeler)
	http.ListenAndServe(":3050", nil)
}
```
<br>

We can do more, using interfaces
``` go
type Info struct {
	Name, Email string
}

func handeler(resw http.ResponseWriter, req *http.Request) {
	info := Info{"Keramot UL Islam", "keramotul.islam@gmail.com"}
	template, _ := template.ParseFiles("./webDev/besics/02.template/public/index.html")
	template.Execute(resw, info)
}
```

then in `./public/index.html`
``` html
<h2>Name {{ .Name }}</h2>
<p>Email {{ .Email }}</p>
```