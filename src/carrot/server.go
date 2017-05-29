package carrot

import (
	"fmt"
	"log"
	"net/http"
)

func renderHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func StartHTTPServer(port string) {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", renderHTML)

	fmt.Printf("HTTP Server Listening at... %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
