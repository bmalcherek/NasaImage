package views

//Package contains funtions which handles endpoints in api

import (
	"NasaImage/db"
	"NasaImage/models"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Images(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")
	var img models.APOD
	filter := bson.D{{"date", "2020-03-14"}}
	err := db.ApodCollection.FindOne(db.Ctx, filter).Decode(&img)
	if err != nil {
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
