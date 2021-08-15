package main

import "net/http"

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello docker"))
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
