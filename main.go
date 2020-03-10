package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

var API_KEY string = "Mv61fomCvGOpcb01LyHq02MziK10cv6tCMxdyfox"

type APOD struct {
	Copyright       string
	Date            string
	Explanation     string
	Hdurl           string
	Media_type      string
	Service_version string
	Title           string
	Url             string
}

func getUrl(url string) APOD {
	resp, err := http.Get(url + "?api_key=" + API_KEY)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	/*body, err := ioutil.ReadAll(resp.Body)*/
	//if err != nil {
	//log.Fatal(err)
	/*}*/

	//var result map[string]interface{}
	var apod APOD
	err = json.NewDecoder(resp.Body).Decode(&apod)
	if err != nil {
		log.Fatal(err)
	}

	return apod
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	image := getUrl("https://api.nasa.gov/planetary/apod")
	_, err = conn.Do("HSET", redis.Args{"APOD:test"}.AddFlat(image)...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
