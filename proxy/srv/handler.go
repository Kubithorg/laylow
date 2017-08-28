package srv

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
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
		query := r.URL.Query()
		fmt.Println(query)
		url := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/hasJoined?username=%v&serverId=%v&ip=%v",
			query.Get("username"),
			query.Get("serverId"),
			query.Get("ip"))
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		response, _ := client.Do(req)
		defer response.Body.Close()
		read, _ := ioutil.ReadAll(response.Body)
		w.Write(read)
	})
}
