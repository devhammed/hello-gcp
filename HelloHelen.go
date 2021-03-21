package helloGCP

import (
        "fmt"
        "net/http"
)

// HelloHelen is an HTTP Cloud Function to greet HelloHelen.
func HelloHelen(w http.ResponseWriter, r *http.Request) {  
        fmt.Fprintf(w, "Hello, Helen. Will you kuku marry me now?")
}

