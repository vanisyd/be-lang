package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/server"
	"web/server/kind"
	"web/server/middleware"
)

func GetResponse(w http.ResponseWriter, r *http.Request, route kind.Route) {
	server.HTTPBody = make(map[string]any)
	server.HTTPRequest = r
	server.HTTPQuery = server.HTTPRequest.URL.Query()
	json.NewDecoder(server.HTTPRequest.Body).Decode(&server.HTTPBody)
	var breakReq bool

	if len(route.Middlewares) > 0 {
		for _, resolver := range route.Middlewares {
			resolver = resolver.Handle(w)
			if resolver.GetStatus() == middleware.MIDDLEWARE_STATUS_BREAK {
				breakReq = true
			}
		}
	}

	if breakReq != true {
		resp := route.Handler()
		w.WriteHeader(resp.StatusCode)
		fmt.Fprintf(w, "%s", resp.Content)
	}
}
