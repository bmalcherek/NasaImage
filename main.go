package main

import (
	"NasaImage/db"
	"NasaImage/handlers"
)

func main() {
	defer db.CtxCancel()
	defer db.Client.Disconnect(db.Ctx)

	handlers.Handle()
}
