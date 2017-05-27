package srv

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Listen() {
	router := httprouter.New()
	router.GET("/hasJoined/:username/:serverId/:ip", hasJoined())
	router.POST("/join", join())
	log.Println("Starting up server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
