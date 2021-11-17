package mathRestApi

import (
	"fmt"
	"log"
	"net/http"
	// "codeblock.ir/helloGoLang/mathOperations"
)

//default page for root address
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Fprintf(w, "\nOpen http://localhost:10000/doc in your browser to see the API documentation")

	fmt.Println("Endpoint Hit: homePage")
}

//show API documentation
func docs(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "http://localhost:10000/ 				      ==> home page")
	fmt.Fprintf(w, "\nhttp://localhost:10000/docs				  ==> api documentation")
	fmt.Fprintf(w, "\nhttp://localhost:10000/sum/{num1}/{num2}    ==> sum 2 numbers")

	fmt.Println("Endpoint Hit: doc")
}

//handle http request
func HandleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/doc", docs)

	fmt.Println("open http://localhost:10000/ in your browser")

	log.Fatal(http.ListenAndServe(":10000", nil))

}
