package main

import (
	"fmt"
	"os"
	"syscall"
    "net/http"
    "log"
)

func main() {

	// Set UID and GID
	_=syscall.Setuid(1000)
	_=syscall.Setgid(1000)

	fmt.Println("Current UID:", os.Getuid())
	fmt.Println("Current GID:", os.Getgid())

    // Serve files from the "public" directory
    fileServer := http.FileServer(http.Dir("./public"))

    // Serve the static site
    http.Handle("/", fileServer)

    // Start the server on port 8080
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}