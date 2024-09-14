package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/tinacious/static-server/utils"
)

func main() {
	port := utils.GetRandomPort()

	// Use configured port
	portString := os.Getenv("PORT")
	if portString != "" {
		i, err := strconv.Atoi(portString)
		if err != nil {
			fmt.Printf("⛔️ invalid port %s - using random port %d\n", portString, port)
		} else {
			port = i
		}
	} else {
		fmt.Printf("🎲 using random port %d - to configure, set the environment variable PORT\n", port)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Printf("🚀 static server at: http://localhost:%d with contents from %s\n", port, cwd)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
