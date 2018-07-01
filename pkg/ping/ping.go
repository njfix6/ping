package main

import (
  "fmt"
  "github.com/njfix6/ping/pkg/request"
  "net/url"
  "time"
)


func Wait(form url.Values) {
  fmt.Println(request.Get("http://localhost:3000/api/jokes"))
  fmt.Println(request.Post("http://localhost:3000/api/jokes/like/:jokeID", form))
  duration := time.Second
  time.Sleep(duration)
}


func main() {
  form := url.Values{
  		"username": {"xiaoming"},
  		"address":  {"beijing"},
  		"subject":  {"Hello"},
  		"from":     {"china"},
  	}
  for i := 1; i <= 1000; i++ {
    go Wait(form)
  }
  fmt.Scanln()

}
