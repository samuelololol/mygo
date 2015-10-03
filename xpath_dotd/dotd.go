package main

import "fmt"
import "net/http"
import "log"
import "github.com/moovweb/gokogiri"
import "io/ioutil"

func main() {

    body := fetch("https://www.manning.com/dotd")
    doc, err := gokogiri.ParseHtml(body)
    if err != nil {
        log.Fatalln(err)
    }
    defer doc.Free()

    html := doc.Root().FirstChild()
    result, err := html.Search("//*[@id=\"top\"]/div[1]/div/div/div/div[2]/div[1]")
    if err != nil {
        log.Fatalln(err)
    }
    //fmt.Println("Hello, playground")
    fmt.Println(result)
    //fmt.Println("Hello, playground")
}

func fetch(url string) []byte {
    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalln(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    return body


}
