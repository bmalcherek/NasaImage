package main

import (
	"encoding/json"
	//"fmt"
	//"io/ioutil"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

var API_KEY string = "Mv61fomCvGOpcb01LyHq02MziK10cv6tCMxdyfox"
var (
	conn redis.Conn
	err  error
)

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

func init() {
	log.Println("Hello from init")
	conn, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
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

func images(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")
	var img APOD
	image, err := redis.Values(conn.Do("HGETALL", "APOD:test"))
	if err != nil {
		log.Fatal(err)
	}
	if err = redis.ScanStruct(image, &img); err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(img)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func main() {
	defer conn.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/images/", images).Methods("GET")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	/* image := getUrl("https://api.nasa.gov/planetary/apod")*/
	//_, err = conn.Do("HSET", redis.Args{"APOD:test"}.AddFlat(image)...)
	//if err != nil {
	//log.Fatal(err)
	//}

	/*fmt.Println()*/
}
