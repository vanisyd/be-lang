package bootstrap

import (
	"net/http"
	"web/controller/router"
	"web/server"
)

func Execute() {
	handler := http.NewServeMux()

	server.BuildRouterMap(router.NRoutes)

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		route := server.GetRouteByPath(path)

		if route.Handler == nil {
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
