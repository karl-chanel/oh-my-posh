package main
import (
    "net/http"
)

func main() {
    // http.Handle("/", http.FileServer(http.Dir("C:\\Users\\devgeek\\Desktop\\pic")))



    
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.ListenAndServe(":8000", nil)
}
