package devhammedGithubIoHelloGCP

import (
	"fmt"
	"net/http"
)

// HelloTomisin is an HTTP Cloud Function to greet Tomisin.
func HelloTomisin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Tomisin. Some people are way too beautiful to be described in words. You’re one of them.")
}
