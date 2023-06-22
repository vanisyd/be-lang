package bootstrap

import (
	"net/http"
	"studying/web/controller/api"
	"studying/web/server"
)

func Execute() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		server.GetResponse(w, r, api.GetWords)
	})

	server.Run()
}
