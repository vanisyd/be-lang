package router

import (
	"web/controller/api"
	"web/server/kind"
	"web/server/middleware"
)

var Routes []kind.Route = []kind.Route{
	{
		Path: "/words",
		Children: []kind.Route{
			{
				Path:    "/",
				Handler: api.GetWords,
				Method:  kind.METHOD_GET,
				Middlewares: []kind.HTTPMiddleware{
					middleware.AuthMiddleware{},
				},
			},
			{
				Path:    "/create",
				Handler: api.AddWord,
				Method:  kind.METHOD_POST,
			},
			{
				Path:    "/update",
				Handler: api.UpdateWord,
				Method:  kind.METHOD_PATCH,
			},
		},
	},
	{
		Path: "/languages",
		Children: []kind.Route{
			{
				Path:    "/",
				Handler: api.GetLangs,
				Method:  kind.METHOD_GET,
			},
			{
				Path:    "/create",
				Handler: api.AddLang,
				Method:  kind.METHOD_POST,
			},
		},
	},
}
