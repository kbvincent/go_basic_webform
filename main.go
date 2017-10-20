package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

type Review struct {
  Name     string
  Rating   string
  Comments string
}

func review(httpWriter http.ResponseWriter, request *http.Request) {
    fmt.Println("method:", request.Method) //get request method
    if request.Method == "GET" {
        form, _ := template.ParseFiles("review.gtpl")
        form.Execute(httpWriter, nil)
    } else {
        request.ParseForm()
        review := Review{
          Name: strings.Join(request.Form["name"],""),
          Rating: strings.Join(request.Form["rating"],""),
          Comments: strings.Join(request.Form["comments"],"")}
        thanks, _ := template.ParseFiles("thanks.gtpl")
        thanks.Execute(httpWriter, review)
    }
}

func main() {
    // Handle Review Page
    http.HandleFunc("/review", review)

    // Handle Static Requests
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    fmt.Println("Serving on port :9090")
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
