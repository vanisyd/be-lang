package router

import (
	"strings"
	"web/server/kind"
)

var routerMap kind.RouterMapping = kind.RouterMapping{}

func GetRouterMap() kind.RouterMapping {
	return routerMap
}

func BuildRouterMap(routes []kind.Route) {
	prepareRouterMap(routes, routerMap)
}

func GetRouteByPath(path string) kind.Route {
	var mappingItem interface{}
	mapping := GetRouterMap()
	pathParts := strings.Split(path, "/")

	for _, part := range pathParts {
		if part != "" {
			if part[:1] != "/" {
				part = "/" + part
			}

			mappingItem = prepareRouterPath(part, mapping, mappingItem)
		}
	}

	route, ok := mappingItem.(kind.RouterMappingItem)
	if ok {
		if route.Route.Handler == nil {
			indexMapping := prepareRouterPath("/", mapping, mappingItem)
			indexRoute, ok := indexMapping.(kind.RouterMappingItem)
			if ok {
				route = indexRoute
			}
		}
		return route.Route
	}

	return kind.Route{}
}

func prepareRouterPath(pathPart string, mapping kind.RouterMapping, mappingItem interface{}) interface{} {
	var curMapping kind.RouterMappingItem
	var ok bool

	curItem, isChild := mappingItem.(kind.RouterMappingItem)
	if !isChild {
		curMapping, ok = mapping[pathPart].(kind.RouterMappingItem)
	} else {
		curMapping, ok = curItem.Children[pathPart].(kind.RouterMappingItem)
	}

	if ok {
		return curMapping
	}

	return nil
}

func prepareRouterMap(routes []kind.Route, mapping kind.RouterMapping) {
	for _, route := range routes {
		mapping[route.Path] = kind.RouterMappingItem{
			Route: kind.Route{
				Path:        route.Path,
				Handler:     route.Handler,
				Method:      route.Method,
				Middlewares: route.Middlewares,
			},
		}
		if route.Children != nil {
			nestedMap := make(kind.RouterMapping)
			prepareRouterMap(route.Children, nestedMap)
			mappingItem, ok := mapping[route.Path].(kind.RouterMappingItem)
			if ok {
				mappingItem.Children = nestedMap
				mapping[route.Path] = mappingItem
			}
		}
	}
}
