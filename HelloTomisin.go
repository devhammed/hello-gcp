package helloGCP

import (
        "fmt"
        "net/http"
)

// HelloTomisin is an HTTP Cloud Function to greet Tomisin.
func HelloTomisin(w http.ResponseWriter, r *http.Request) {  
        fmt.Fprintf(w, "Hello, Tomisin. Our very own biggest head!")
}
