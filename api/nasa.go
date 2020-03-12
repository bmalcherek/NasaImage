package api

// All interactions with NASA api here

import (
	"encoding/json"
	"log"
	"net/http"

	"NasaImage/models"
)

var API_KEY string = "Mv61fomCvGOpcb01LyHq02MziK10cv6tCMxdyfox"

func GetUrl(url string) models.APOD {
	resp, err := http.Get(url + "?api_key=" + API_KEY)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var apod models.APOD
	err = json.NewDecoder(resp.Body).Decode(&apod)
	if err != nil {
		log.Fatal(err)
	}

	return apod
}
