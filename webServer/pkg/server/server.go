package server

import "net/http"

//server is an HTTP server
type Server struct{

}

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


//HandleInder handles the index path "/".
func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexPage))
}

//HandleInder handles the users path "/users".
func (s *Server) HandleUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userInfo))
}