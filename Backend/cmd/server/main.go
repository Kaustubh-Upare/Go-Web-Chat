package main

import (
	"kaustubh-upare/chatApplication/internals/server"
	"net/http"
)

func main() {

	server.SetupRoutes()

	// We do nil after 8000 bcz we are doing websocket
	http.ListenAndServe(":9000", nil)
}
