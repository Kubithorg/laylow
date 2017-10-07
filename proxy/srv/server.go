package srv

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Listen() {
	router := httprouter.New()
	router.GET("/hasJoined", hasJoined())
	router.GET("/users/profiles/minecraft", accountInfo())
	log.Println("Starting up server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
