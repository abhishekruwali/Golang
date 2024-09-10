package main

import (
	"fmt"
	"net/url"
)

func main() {

	// string to url form convert
	fmt.Println("learning Url..")
	myURL := "https://example.com:8080/path/toresource?key1=value1&key2=value2"
	fmt.Printf("Type of URL %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("cant parse url", err)
		return
	}
	fmt.Printf("Type of URL : %T\n", parsedURL)
	fmt.Println("Scheme", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("path:", parsedURL.Path)
	fmt.Println("raw", parsedURL.RawQuery)

	// Modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=iamprince"

	// constructing a URL string from a Url object
	newUrl := parsedURL.String()

	fmt.Println("new URL :", newUrl)
	fmt.Printf("new URL %T", newUrl)

}
