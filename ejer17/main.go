//Servidor Web en GO
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3010", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
