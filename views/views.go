package views

//Package contains funtions which handles endpoints in api

import (
	"NasaImage/db"
	"NasaImage/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ImagesList(w http.ResponseWriter, r *http.Request) {
	var results []*models.APOD

	findOptions := options.Find()
	findOptions.SetLimit(20)

	cur, err := db.ApodCollection.Find(db.Ctx, bson.D{{}}, findOptions)
	defer cur.Close(db.Ctx)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(db.Ctx) {
		var picture models.APOD
		err = cur.Decode(&picture)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &picture)
	}

	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Images(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var img models.APOD

	filter := bson.D{{"date", vars["date"]}}
	log.Println("Date var: ", vars["date"])
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
