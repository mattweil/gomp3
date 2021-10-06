package main

import (
   "io/ioutil"
   "log"
   "net/http"
)

func main() {
	URL := ""
   	resp, err := http.Get(URL)
   	if err != nil {
   	   log.Fatalln(err)
   	}

   	body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
   	   log.Fatalln(err)
   	}

   	sb := string(body)
   	log.Printf(sb)
}