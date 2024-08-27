package main

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
/ - root
/home
/support
*/

var htmltag = `<html>
<title> demo page</title>
<body bgcolor = red>
<h1>
this is our demo <b>response</b>
</h1>
</body>
</html>`
var (
	rootCount    = 0
	homeCount    = 0
	supportCount = 0
)

func main() {
	fmt.Println("Demo: http service")
	//rootCount := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootCount++
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Add("ContentLenght", strconv.Itoa(len(htmltag)))
		w.Header().Add("ContentType", "text/html")
		w.Write([]byte(htmltag))
		w.WriteHeader(http.StatusOK)
	})
	//homeCount := 0
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		homeCount++
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("You are in home.."))
		w.WriteHeader(http.StatusOK)
	})
	//supportCount := 0
	http.HandleFunc("/support", func(w http.ResponseWriter, r *http.Request) {
		supportCount++
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("You are in support.."))
		w.WriteHeader(http.StatusOK)
	})
	go ReportingService()
	err := http.ListenAndServe(":2020", nil)
	if err != nil {
		fmt.Println("Received err", err)
	}

}

func ReportingService() {

	var handle = http.NewServeMux()
	handle.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintln("RootCount", rootCount, "homeCount", homeCount, "support count", supportCount)))
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":3030", handle)
}
