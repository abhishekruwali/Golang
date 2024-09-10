package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"comp"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(" Error getting ", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response", res.Status)
		return
	}
	// json
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(" Error getting ", err)
	// 	return
	// }
	// fmt.Printf("DATA %T \n:", data)

	// json convert todo type
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo) // data -> normal object -> save todo
	if err != nil {
		fmt.Println("Error decoding :", err)
		return
	}
	fmt.Println("Todo:", todo)
	fmt.Println("Todo: ", todo.Title)
	fmt.Println("completed response :", todo.Completed)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "PrinceKumar",
		Completed: true,
	}
	//    Convert the Todo struct to JSON server except
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	// convert json to string it should be readable
	jsonString := string(jsonData)

	strings.NewReader(jsonString)

	//    This function creates a new Reader that reads from the
	//    provided string (jsonString).
	//    The Reader implements the io.Reader interface,
	//    which allows you to pass it to functions like http.Post().
	//  convert to reader from string
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println(" type of data", data)
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    23789,
		Title:     "Prince Kumar Golang",
		Completed: true,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	//  create PUT request

	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	//  Content type
	req.Header.Set("Content-type", "application/json")

	// Make a client

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))

	// Send the Request

}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	// Create Delete Request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating Delete Request", err)
		return
	}

	// send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status :", res.Status)
}

func main() {
	fmt.Println("Learning crud ...")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()

}
