package devhammedGithubIoHelloGCP

import (
	"fmt"
	"net/http"
)

// HelloAzeezat is an HTTP Cloud Function to greet Azeezat.
func HelloAzeezat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Azeezat. My very own sweetheart!")
}
