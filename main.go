// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	fmt.Printf("Start server");

// 	routes := httprouter.new()
// 	server := http.server{Addr: "localhost:8888", Handler: routes}
// }

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %q", html.EscapeString((r.URL.Path)))
	})

	fmt.Println("peepee poopoo")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
