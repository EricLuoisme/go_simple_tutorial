package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	get, err := http.Get("http://www.google.com/robots.txt")
	if nil != err {
		log.Fatal(err)
	}
	// call defer first, before further using
	defer get.Body.Close()

	robots, err := ioutil.ReadAll(get.Body)
	if nil != err {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}
