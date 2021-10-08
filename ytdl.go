package main

import (
    "fmt"
    "os"
    "io/ioutil"
   	"log"
   	"net/http"
)

func main() {


    if (len(os.Args ) != 2) {
    	fmt.Println("usage: ytdl.go url")
    	return
    }


    url := os.Args[1]

    resp, err := http.Get(url)
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

