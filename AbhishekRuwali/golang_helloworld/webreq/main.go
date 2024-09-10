package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response,err")
		return
	}
	defer res.Body.Close() // network connection resources use
	fmt.Printf("Type of response %T\n", res)

	// read the response body
	data, err := ioutil.ReadAll(res.Body)
	// readAll array of bytes return
	// data is in byte format
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("response ", string(data))

}
