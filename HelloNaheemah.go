package devhammedHelloGCP

import (
	"fmt"
	"net/http"
)

// HelloNaheemah is an HTTP Cloud Function to greet Naheemah.
func HelloNaheemah(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Naheemah. My sweetheart, I must be the luckiest man to have someone like you in my life! Every day of my life is beautiful because of you! You are my everything!")
}
