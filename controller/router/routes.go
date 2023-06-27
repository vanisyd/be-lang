package router

import (
	"web/controller/api"
	"web/server"
)

var Routes map[string]server.Route = map[string]server.Route{
	"/words/": {
		Path:    "/words",
		Handler: api.GetWords,
		Method:  server.METHOD_GET,
	},
	"/words/create/": {
		Path:    "/words/create",
		Handler: api.AddWord,
		Method:  server.METHOD_POST,
	},
	"/words/update/": {
		Path:    "/words/update",
		Handler: api.UpdateWord,
		Method:  server.METHOD_PATCH,
	},
	"/languages/": {
		Path:    "/languages",
		Handler: api.GetLangs,
		Method:  server.METHOD_GET,
	},
	"/languages/create/": {
		Path:    "/languages/create",
		Handler: api.AddLang,
		Method:  server.METHOD_POST,
	},
}
