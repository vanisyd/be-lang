package bootstrap

import (
	"net/http"
	"studying/web/controller/router"
	"studying/web/server"
)

func Execute() {
	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path[len(path)-1:] != "/" {
			path += "/"
		}

		route, ok := router.Routes[path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if route.Method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		server.GetResponse(w, r, route.Handler)
	})

	server.Run(handler)
}
