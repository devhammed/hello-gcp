package devhammedGithubIoHelloGCP

import (
	"fmt"
	"net/http"
)

// HelloRidwan is an HTTP Cloud Function to greet Ridwan.
func HelloRidwan(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Ridwan. Your support and encouragement have always been the driving factors in my life. I want you to know that I appreciate your cordiality with all my life!")
}
