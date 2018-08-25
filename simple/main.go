package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url string
)

//HttpClient allow dep inj w/interface
type HttpClient interface {
	Get(string) (*http.Response, error)
}

func init() {
	flag.StringVar(&url, "url", "http://google.com", "what url to parse?")
	flag.Parse()
}

func main() {
	client := &http.Client{}
	err := send(client, url)
	if err != nil {
		panic(err)
	}

}
func send(client HttpClient, link string) error {

	response, err := client.Get(link)
	if err != nil {
		return err
	}
	if response == nil {
		fmt.Println("empty response")
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("can't read body")
		return err
	}
	fmt.Println(string(body))
	return nil

}
