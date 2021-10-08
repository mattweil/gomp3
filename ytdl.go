package main

import (
    "fmt"
    "os"
    "io/ioutil"
   	"log"
   	"net/http"
   	"regexp"
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

   	r, _ := regexp.Compile("https:\\/\\/r[0-9]---[a-z][a-z]-[a-z0-9]{8}\\.googlevideo\\.com")
   	
   	fmt.Println(r.FindAllString(sb, -1))

}

