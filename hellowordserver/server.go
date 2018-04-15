package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.FormValue("name"))
		fmt.Fprintf(writer,"<h1>hello word %s !<h1>", request.FormValue("name"))
	})

	http.ListenAndServe(":8888", nil)
}
