//make sure to break your code in multiple functions and when it makes sense, break it in different packages and files

package main

import (
	"fmt"
	"net/http"
	"os"
)

const name = "Paulo \"The King\""

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	message := fmt.Sprintf("Hello %s\n", name)
	res.WriteHeader(http.StatusOK) //Status code 200
	res.Write([]byte(message))
}
