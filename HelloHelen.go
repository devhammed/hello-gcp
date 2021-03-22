package devhammedHelloGCP

import (
	"fmt"
	"net/http"
)

// HelloHelen is an HTTP Cloud Function to greet Helen.
func HelloHelen(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Helen. I can’t imagine a life without you, darling. I feel lucky to have you in my life.")
}
