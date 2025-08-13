package main

import "net/http"

func serverWs(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// We do nil after 8000 bcz we are doing websocket
	http.ListenAndServe(":8000", nil)
}
