package main


import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})



	//When running on Bluemix, get the PORT from the environment variable.
	port := os.Getenv("PORT")
	if port == "" {
		port = "13100" //Local
	}
	//r.Run(":" + port)
	http.ListenAndServe(":" + port, nil)
}
