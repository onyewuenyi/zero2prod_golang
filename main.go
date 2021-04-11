package main

import (
	"log"
	"net/http"
)

// Sign up a new subscriber: As the blog author,I want to send an email
// to all my subscribers,So that I can notify them when new content is published
// 	   1. Simple http server that can handle http req
//     2. /health_check GET 200 OK res
//         Define route/URL path and guards
//         Handler fun to handle req to our root URL/endpoint/route
//         Unit test
//     3. Subscriptions endpoint and handler
//         Req data
//         Handler
//         Setup DB
//         Unit test: endpoint and DB

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func main() {
	// http://localhost:10000/health_check
	http.HandleFunc("/health_check", healthCheck)

	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// ⚡⚡⚡ Setup Cheat Sheet ⚡⚡⚡
// cd $HOMEPATH
// git clone YOUR_REPO_URL
// cd YOUR_REPO_DIRECTORY
// go mod init github.com/<your GitHub username>/<project name>
// go build
// go run .
