package main

import (
	"errors"
	"fmt"
	"main/controllers"
	"net/http"
)

// type HTTPRequest struct {
// 	Method string
// }

func main() {
	// port := 3000
	// fmt.Println("Hello Gohpers!")
	// _, err := StartWebServer(port)
	// fmt.Println(err)
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	// example switch
	// r := HTTPRequest{Method: "GET"}

	// switch r.Method {
	// case "GET":
	// 	println("GET request")
	// case "POST":
	// 	println("POST request")
	// case "PUT":
	// 	println("PUT request")
	// default:
	// 	print("Unhandled method")
	// }

}

func StartWebServer(port int) (int, error) {
	fmt.Println("Server starting on port: ", port)
	// do something here
	// panic("Something bad just happened")
	fmt.Println("Server started on port: ", port)
	return port, errors.New("an error has occurred")
}
