package main

import "net/http"

func main() {
	http.HandlerFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}
