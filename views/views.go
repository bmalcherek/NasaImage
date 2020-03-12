package views

//Package contains funtions which handles endpoints in api

import (
	"encoding/json"
	"log"
	"net/http"

	"NasaImage/db"
	"NasaImage/models"

	"github.com/gomodule/redigo/redis"
)

func Images(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")
	var img models.APOD
	image, err := redis.Values(db.Conn.Do("HGETALL", "APOD:test"))
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

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
