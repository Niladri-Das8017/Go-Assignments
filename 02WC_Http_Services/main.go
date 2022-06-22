/*WC_HttpServics:	A small service that accepts as input a body of text, such as that from a book,
and returns the top ten most-used words along with how many times they occured in the text. */

package main

import (
	"fmt"
	"httpservices/routers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("\n******Welcome to Http Services*****")

	//Routing
	r := routers.Router()
	fmt.Println("Server is getting started...")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4040", r))

}
