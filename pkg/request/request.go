package request

import (
  "net/http"
  "strings"
  "net/url"
)


func Get(url string) (resp *http.Response, err error){
  return http.Get(url)
}

func Post(url string, form url.Values) (resp *http.Response, err error){
  return  http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
}

func Expect() {
  
}
