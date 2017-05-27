package srv

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type JoinRequest struct {
	AccessToken     string `json:"accessToken"`
	SelectedProfile string `json:"selectedProfile"`
	ServerId        string `json:"serverId"`
}

func hasJoined() httprouter.Handle {
	client := http.Client{}
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		url := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/hasJoined?username=%v&serverId=%v&ip=%v",
			params.ByName("username"),
			params.ByName("serverId"),
			params.ByName("ip"))
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		response, _ := client.Do(req)
		defer response.Body.Close()
		read, _ := ioutil.ReadAll(response.Body)
		w.Write(read)
	})
}

func join() httprouter.Handle {
	client := http.Client{}
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if content, ok := readOrFail(w, r.Body); ok {
			req, _ := http.NewRequest(http.MethodPost, "https://sessionserver.mojang.com/session/minecraft/join", nil)
			req.Header.Set("Content-Type", "application/json")
			response, _ := client.Do(req)
			defer response.Body.Close()
			content, _ := ioutil.ReadAll(response.Body)
			w.Write(content)
		}
	})
}

// readOrFail reads the body and if an error has been encountered, it writes the error to the http
// ResponseWriter and returns and an uninitialized array and false; otherwise, returns the body and
// true.
func readOrFail(w http.ResponseWriter, body io.ReadCloser) ([]byte, bool) {
	var ret []byte
	if ret, err := ioutil.ReadAll(body); err != nil {
		log.Printf("Could not ready body: %v", err)
		http.Error(w, "cannot read body", http.StatusBadRequest)
		return ret, false
	}
	return ret, true
}
