package mathRestApi

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {

	http.HandleFunc("/", homePage)

	fmt.Println("open localhost::10000 in your browser")

	log.Fatal(http.ListenAndServe(":10000", nil))

}
