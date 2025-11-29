package main
import ("fmt"; "net/http"; "strings")
func handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" { fmt.Fprintf(w, "Welcome to root web") } else { fmt.Fprintf(w, "you access the %s", path) }
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}