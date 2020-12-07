# Create a server

Easiest basic way to create a server and listen on a port.
``` go
func handeler(resw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resw, "Hello Universe")
}	

func main() {
	http.HandleFunc("/", handeler)
	http.ListenAndServe(":3050", nil)
}
```	
<br>

Let's add a 404 route, and set `content-type` to `text/html`
``` go
func handeler(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "text/html")

	if req.URL.Path == "/" {
		fmt.Fprint(resw, "<h1>Hello Universe</h1>")
	} else {
		resw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resw, "<h1>404 Nothing Found</h1>")
	}
}	
```
