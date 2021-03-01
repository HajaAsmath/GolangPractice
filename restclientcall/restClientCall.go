package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://api.theysaidso.com/qod")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(resp))
	http.Handle("/sayhello", new(handle))
	http.ListenAndServe(":8080", nil)
}

type handle struct {
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.PostFormValue("name"))
	str := "<html>" + r.URL.Path
	fmt.Fprintf(w, "Hello Welcome %q", html.EscapeString(str))
}
