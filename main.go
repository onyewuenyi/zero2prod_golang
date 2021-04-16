package main

import "os"

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

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PW"),
		os.Getenv("APP_DB_NAME"))

	app.Run(":8010")
}

// ⚡⚡⚡ Setup Cheat Sheet ⚡⚡⚡
// cd $HOMEPATH
// git clone YOUR_REPO_URL
// cd YOUR_REPO_DIRECTORY
// go mod init github.com/<your GitHub username>/<project name>
// go build
// go run .
