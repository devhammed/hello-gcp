package devhammedGithubIoHelloGCP

import (
	"fmt"
	"net/http"
)

// HelloGbenga is an HTTP Cloud Function to greet Gbenga.
func HelloGbenga(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Gbenga. Working with a superlative boss like is the best experience of my career. Take my heartiest thanks for your excellent leadership.")
}
