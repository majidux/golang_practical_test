package main

import "net/http"

func SimpleServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello go"))
	})
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		panic(err.Error())
	}
}
