package webserver

import (
	"log"
	"net/http"
	"time"
)

var indexPage = `<!DOCTYPE html>
<html>
	<body>
		<h1>My First Heading </h1>
		<p>My First paragraph. </p>
	</body>
</html>
`

var userInfo = `{
	"name": "testuser",
	"age": 21
}`

func WebServer() {
	address := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(indexPage))
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(userInfo))
	})

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Start server : %v", address)
	log.Fatal(s.ListenAndServe())
}