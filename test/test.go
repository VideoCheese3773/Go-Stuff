// Eduardo Mejia
package main

import (
	"fmt"      // format for printing and scan the URL
	"io"       // read returned info from api
	"net/http" // this is to make the API call
	"os"       // for the exit code to use as a break
)

//import "os"

func TLSSecurityCheck(url string) {
	fmt.Println("URL:", url)

	resp, err := http.Get(url)
	if err != nil { // This is to make sure there are no errors in getting stuff from the API, nil is the go version of null btw
		fmt.Println(err.Error()) //If there is an error, this will print it
		os.Exit(1)               //A number means there was an error, 0 is success
	}
	respData, err := io.ReadAll(resp.Body)
	if err != nil { // Same thing as the one above, just making sure there are no errors in the information recived this time
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(respData)) //needs string to be readable, otherwise its just numbers
}

func main() {
	var url string
	fmt.Println("Please write the URL to check their TLS Security")
	fmt.Scan(&url) // This is how you save inputs from the console to a variable
	if url == "a" {
		url = "http://pokeapi.co/api/v2/pokedex/kanto/" // This whole if section is just a placeholder url to test if the API call works
	}

	TLSSecurityCheck(url)
}
