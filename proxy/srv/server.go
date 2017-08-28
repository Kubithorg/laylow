package srv

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Listen() {
	router := httprouter.New()
	router.GET("/hasJoined", hasJoined())
	log.Println("Starting up server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
