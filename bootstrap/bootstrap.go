package bootstrap

import (
	"log"
	"net/http"
	apiRouter "web/controller/router"
	"web/server"
	"web/server/http/response"
	"web/server/http/router"
)

func Execute() {
	Init()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	handler := http.NewServeMux()

	router.BuildRouterMap(apiRouter.Routes)

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		server.CurrentRoute = router.GetRouteByPath(path)

		if server.CurrentRoute.Handler == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if server.CurrentRoute.Method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		response.GetResponse(w, r, server.CurrentRoute)
	})

	server.Run(handler)
}
